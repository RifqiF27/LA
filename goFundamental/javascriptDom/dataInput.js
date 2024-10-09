let dataArray = [];

document.getElementById("saveButton").addEventListener("click", function () {
  const nameInput = document.getElementById("name").value;
  if (nameInput) {
    dataArray.push(nameInput);
    document.getElementById("name").value = "";
    alert("Data saved successfully!");
  } else {
    alert("Please enter a name.");
  }
});

function showDataList() {
  const dataListDiv = document.getElementById("dataList");
  const listElement = document.getElementById("list");

  listElement.innerHTML = "";

  if (dataArray.length > 0) {
    for (let i = 0; i < dataArray.length; i++) {
      const li = document.createElement("li");
      li.textContent = dataArray[i];
      listElement.appendChild(li);
    }
    dataListDiv.style.display = "block";
  } else {
    dataListDiv.style.display = "block";
    listElement.innerHTML = "<h4>No data available.</h4>";
  }

  document.getElementById("container").style.display = "none";
}

function showInputForm() {
  document.getElementById("dataList").style.display = "none";
  document.getElementById("container").style.display = "block";
}
