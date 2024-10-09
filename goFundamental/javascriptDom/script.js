let product = [
  {
    nama: "Baju",
    stock: 10,
    price: 300000,
    description: "Baju Lumushive",
  },
  {
    nama: "Tas",
    stock: 5,
    price: 100000,
    description: "Tas Lumushive",
  },
  {
    nama: "Celana",
    stock: 7,
    price: 200000,
    description: "Celana Lumushive",
  },
];

function search() {
  for (let i = 0; i < product.length; i++) {
    const element = product[i];
    if (element.nama == "Tas") {
      // console.log(element);
      return { nama: element.nama, stock: element.stock, price: element.price };
    }
  }
  return "not found";
}
console.log(search());

const text = document.querySelector(".text");
const button = document.getElementById("btn");

const originalText = text.innerHTML;
let count = 0;

button.addEventListener("click", function () {
    count++
  if (count == 1) {
    text.innerText = "Aku Programmer handal";
  } else if (count == 2) {
    text.innerText = "Aku pasti bisa jadi programmer";
  } else if (count == 3) {
    text.innerText = originalText;
    count = 0;
  }
});

document.getElementById("btn").addEventListener("click", function () {
  let list = document.getElementById("list");
  let newItem = document.createElement("li");
  newItem.innerHTML = document.getElementById("in").value;

  let removeButton = document.createElement("span");
  removeButton.innerHTML = " ⨉"; 
  removeButton.style.cursor = "pointer"; 
  removeButton.style.color = "red"; 

  removeButton.addEventListener("click", function () {
    list.removeChild(newItem);
  });

  newItem.appendChild(removeButton);
  list.appendChild(newItem);
});


