function loadData() {
  fetch(
    "https://lumoshive-academy-media-api.vercel.app/api/games?page=2&search",
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  )
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok", error);
      }
      return response.json();
    })
    .then((data) => {
      let content = "";

      for (let i = 0; i < data.length; i++) {
        const element = data[i];

        content += `              
                    <h3>No ${i + 1}</h3>
                    <img src="${element.thumb}"> 
                    <h3>${element.title}</h3>
                    <p>${element.desc}</p>
                    <hr>    
            `;
      }

      document.getElementById("content").innerHTML = content;
    })
    .catch((error) => {
      console.error("There was a problem with the fetch operation:", error);
    });
}
function listUser() {
  fetch("https://reqres.in/api/users?page=2", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok", error);
      }
      return response.json();
    })
    .then((data) => {
      console.log(data.data);

      let content = "";

      for (let i = 0; i < data.data.length; i++) {
        const element = data.data[i];
        console.log(element);

        content += `              
                    <p>No ${i + 1}</p>                 
                    <p>email : ${element.email}</p>
                    <p>First Name : ${element.first_name}</p>
                    <p>Last Name : ${element.last_name}</p>
                    <hr>    
            `;
      }

      document.getElementById("listUser").innerHTML = content;
    })
    .catch((error) => {
      console.error("There was a problem with the fetch operation:", error);
    });
}
function singleUser(id) {
  fetch("https://reqres.in/api/users/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok", error);
      }
      return response.json();
    })
    .then((data) => {
      console.log(data);

      let content = `              
                    
                    <img src="${data.data.avatar}">
                    <p>email : ${data.data.email}</p>
                    <p>First Name : ${data.data.first_name}</p>
                    <p>Last Name : ${data.data.last_name}</p>
                    <p>url: ${data.support.url}</p>
                    <p>text: ${data.support.text}</p>
                    <hr>    
            `;

      document.getElementById("singleUser").innerHTML = content;
    })
    .catch((error) => {
      console.error("There was a problem with the fetch operation:", error);
    });
}
