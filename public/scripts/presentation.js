/** Generates HTML based upon an object of type {our html type}. */
const generateElement = (obj) => {
  // If not a good JSON, dont do anything
  if (!obj || !obj.type) return;

  // Create the DOM element as per type attribute
  let element = document.createElement(obj.type),
    prop;

  // For all properties in the JSON
  for (prop in obj) {
    if (obj.hasOwnProperty(prop)) {
      // If it is type, no need to handle, already done. Skip and continue.
      if (prop === "type") continue;

      // It is content, create the content
      if (prop === "content") {
        // If the value is string or number, create a text node
        if (typeof obj[prop] === "string" || typeof obj[prop] === "number") {
          element.appendChild(document.createTextNode(obj[prop]));
        }
        // If it is a list, iterate and handle recursively
        else if (Array.isArray(obj[prop])) {
          var tempArray = obj[prop];
          var i = 0,
            l = tempArray.length;
          for (; i < l; i++) {
            // Fixed for a Node appendChild error
            if (typeof tempArray[i] === "string") {
              element.innerHTML += tempArray[i];
            } else {
              element.appendChild(generateElement(tempArray[i]));
            }
          }
        }
        // Otherwise its an object, handle recursively
        else {
          element.appendChild(generateElement(obj[prop]));
        }
      }
      // Otherwise it is an attribute, add the attribute
      else {
        element.setAttribute(prop, obj[prop]);
      }
    }
  }
  return element;
};

const getPresentation = async () => {
  const response = await fetch("http://10.0.11.84:8000/device/presentation", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  });
  return await response.json();
};

let previousTemplate;

const updateCurrentPresentation = async () => {
  try {
    const presentation = await getPresentation();
    const template = presentation.pages[0].component;
    if (JSON.stringify(previousTemplate) === JSON.stringify(template)) return;
    previousTemplate = template;
    document.title = `E-hall | ${presentation.name}`;
    document.body.innerHTML = "";
    document.body.appendChild(generateElement(template));
  } catch (e) {
    console.error(e);
    document.getElementById("error-box").style.display = "block";
  }
};

updateCurrentPresentation();

setInterval(async () => {
  await updateCurrentPresentation();
}, 1000);

// template = {
//   type: "div",
//   style: "height: 100vh;width: 100vw;background-color:#7f997e",
//   content: [
//     {
//       type: "h1",
//       content: "Título 1",
//       style: "font-weight:bold",
//     },
//     {
//       type: "p",
//       content: "texto normal",
//       style: "color:#f1f1f1;",
//     },
//     {
//       type: "div",
//       content: {
//         type: "p",
//         content: "",
//         style: "position:absolute;",
//       },
//       style:
//         "height:100px;width:100px;border-radius:100%;background-color:grey;position:relative;",
//     },
//   ],
// };
