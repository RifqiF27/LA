const contentData = {
  skills: [
    "Node.js",
    "Python",
    "Golang",
    "GORM",
    "PostgreSQL",
    "Gin",
    "GraphQL",
    "Vue",
    "Redis",
    "React",
    "React Native",
    "Next",
    "AWS",
    "Git",
  ],
  education: [
    "Lumoshive Academy - Backend Golang",
    "Polytechnic State Of Jakarta - Electrical Engineering",
  ],
  experience: [
    "PT Suzuki Indomobil Motor - Technical Control",
    "PT Kaldu Sari Nabati - Maintenance",
  ],
  certification: [
    "Lumoshive Academy - Backend Golang",
    "HackerRank - Javascript intermediate",
  ],
};

function click(type) {
  let content = "";

  const items = contentData[type];
  console.log(type, "ini items");

  if (type === "skills") {
    content += `<ul class="${type}">`;
    for (let i = 0; i < items.length; i++) {
      content += `<li>${items[i]}</li>`;
    }
    content += `</ul>`;
  } else {
    content += `<ul class="content">`;
    for (let i = 0; i < items.length; i++) {
      content += `<li>${items[i]}</li>`;
    }
    content += `</ul>`;
  }

  document.querySelector(".aboutSkillsList").innerHTML = content;
}

document.querySelectorAll(".aboutLinkText").forEach((link) => {
  link.addEventListener("click", function () {
    const target = this.getAttribute("data-target");

    document.querySelectorAll(".aboutLinkText").forEach((link) => {
      link.classList.remove("active");
    });

    this.classList.add("active");

    click(target);
  });
});

document.addEventListener("DOMContentLoaded", function () {
  click("skills");
  document
    .querySelector('.aboutLinkText[data-target="skills"]')
    .classList.add("active");
});

document
  .getElementById("formContact")
  .addEventListener("submit", function (event) {
    event.preventDefault();
    let subject = document.getElementById("subject").value;
    let message = document.getElementById("message").value;

    window.location.href = `mailto:rifqifadluloh27@gmail.com?subject=${subject}&body=${message} text`;

    setTimeout(function () {
      window.location.reload();
    }, 1000);
  });
