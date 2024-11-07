async function fetchProtectedData() {
  const token = localStorage.getItem("authToken");
  console.log("Sending token:", token);

  if (!token) {
    console.log("No token found, please login.");
    alert("You must log in first!");
    return;
  }
  try {
    const response = await fetch("http://localhost:8080/dashboard", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    console.log("Authorization header:", `Bearer ${token}`); // Log header

    if (!response.ok) {
      console.error("Unauthorized:", response.status);
      throw new Error("Unauthorized. Please log in again.");
    }

    const data = await response.json();
    console.log("Protected data:", data);
  } catch (error) {
    console.error("Error fetching protected data:", error);
  }
}

window.addEventListener("load", function () {
  fetchProtectedData();
});
