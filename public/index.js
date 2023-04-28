const loginForm = document.getElementById("login-form");
const emailInput = document.getElementById("email");
emailInput.focus();
const deviceInput = document.getElementById("device");
const passwordInput = document.getElementById("password");
const errorBox = document.getElementById("error-box");
const errorBoxContent = document.getElementById("error-box-content");
const submitButton = document.getElementById("submit-button");

/** Returns a boolean telling whether the given email is valid or not. */
const validateEmail = (email) => {
  return email
    .toString()
    .toLowerCase()
    .match(
      /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    );
};

/** TODO: Implement this method in case it's needed or remove it otherwise. */
const validateDevice = (_device) => {
  return true;
};

/** TODO: Implement this method in case it's needed or remove it otherwise. */
const validatePassword = (_password) => {
  return true;
};

const showErrorBox = (message) => {
  errorBox.style.display = "block";
  errorBoxContent.innerText = message;
};

const hideErrorBox = () => {
  errorBox.style.display = "none";
  errorBoxContent.innerText = "";
};

const startLoading = () => {
  submitButton.innerHTML = `<div class="lds-ring"><div></div><div></div><div></div><div></div></div>`;
  submitButton.disabled = true;
  submitButton.style.cursor = "wait";
};

const stopLoading = () => {
  submitButton.innerHTML = "Acessar apresentação";
  submitButton.disabled = false;
  submitButton.style.cursor = "pointer";
};

const login = async () => {
  const isValidEmail = validateEmail(emailInput.value);
  if (!isValidEmail) {
    showErrorBox("E-mail é inválido.");
    return;
  }

  const isValidDevice = validateDevice(deviceInput.value);
  if (!isValidDevice) {
    showErrorBox("Dispositivo é inválido.");
    return;
  }

  const isValidPassword = validatePassword(passwordInput.value);
  if (!isValidPassword) {
    showErrorBox("Senha é inválida.");
    return;
  }

  const response = await fetch("http://10.0.11.135:8000/device/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    referrerPolicy: "no-referrer",
    body: JSON.stringify({
      email: emailInput.value,
      deviceName: deviceInput.value,
      password: passwordInput.value,
    }),
  });
  if (response.status === 200) {
    const token = await response.text();
    localStorage.setItem("token", token);
    window.location.href = "./pages/presentation.html";
  } else {
    throw new Error();
  }
};

loginForm.onsubmit = async (event) => {
  event.preventDefault();
  event.stopPropagation();
  try {
    hideErrorBox();
    startLoading();
    await login(emailInput.value, deviceInput.value, passwordInput.value);
  } catch (e) {
    showErrorBox("Erro ao fazer login. Tente novamente.");
  } finally {
    stopLoading();
  }
};

emailInput.oninput = (event) => {
  event.preventDefault();
  hideErrorBox();
};

deviceInput.oninput = (event) => {
  event.preventDefault();
  hideErrorBox();
};

passwordInput.oninput = (event) => {
  event.preventDefault();
  hideErrorBox();
};
