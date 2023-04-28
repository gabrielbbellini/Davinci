// TODO: DELETE THIS FILE AND (IT'S MEANT TO BE USED ONLY FOR TESTS)

const renderElement = (obj) => {
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
              element.appendChild(renderElement(tempArray[i]));
            }
          }
        }
        // Otherwise its an object, handle recursively
        else {
          element.appendChild(renderElement(obj[prop]));
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

const testJSon = {
  type: "div",
  content: {
    type: "p",
    content: {
      type: "h2",
      content: "Título 2",
    },
  },
};

const button = document.getElementById("render-template-button");
const textarea = document.getElementById("textarea");
const templateSlot = document.getElementById("template-slot");

button.addEventListener("click", async (event) => {
  event.preventDefault();
  templateSlot.appendChild(renderElement(JSON.parse(textarea.value)));
});
