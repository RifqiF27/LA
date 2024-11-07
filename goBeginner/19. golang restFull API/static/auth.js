async function login(username, password) {
  try {
    const response = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
      throw new Error("Login failed. Please check your credentials.");
    }

    const data = await response.json();
    console.log(data);

    if (data.token) {
      localStorage.setItem("authToken", data.token);
      console.log("Login successful. Token stored in localStorage.", data.token);
      alert("Login successful!");
      window.location.href = "/dashboard";
    } else {
      throw new Error("Invalid response from server.");
    }
  } catch (error) {
    console.error("Error:", error);
    alert(error.message);
  }
}



document
  .getElementById("loginForm")
  .addEventListener("submit", function (event) {
    event.preventDefault();
    
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    if (!username || !password) {
      alert("Please enter both username and password.");
      return;
    }

    login(username, password);
  });
