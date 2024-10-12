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
  let content = ""; // Variable untuk menyimpan konten HTML
    
  // Ambil array yang sesuai dari contentData berdasarkan type
  const items = contentData[type];
  console.log(type,"ini items");
  

  // Loop melalui array dan buat elemen <li> untuk setiap item
  if (type === "skills") {
    content += `<ul class="${type}">`; // Open <ul> for skills
    for (let i = 0; i < items.length; i++) {
      content += `<li>${items[i]}</li>`; // Add each skill as <li>
    }
    content += `</ul>`; // Close <ul> for skills
  } else {
    content += `<ul class="content">`; // Open <ul> for other categories
    for (let i = 0; i < items.length; i++) {
      content += `<li>${items[i]}</li>`; // Add each item as <li>
    }
    content += `</ul>`; // Close <ul> for other categories
  }

  // Masukkan konten <li> ke dalam elemen dengan class .content
  document.querySelector(".AboutSkillsList").innerHTML = content;
}

document.querySelectorAll(".AboutLinkText").forEach((link) => {
  link.addEventListener("click", function () {
    const target = this.getAttribute("data-target"); // Ambil target dari atribut data-target
    // console.log(target, "ini target");
    // Remove active class from all links
    document.querySelectorAll(".AboutLinkText").forEach((link) => {
        link.classList.remove("active");
      });
  
      // Add active class to the clicked link
      this.classList.add("active");
  
    
    click(target); // Panggil fungsi showContent dengan target yang sesuai
  });
});

document.addEventListener("DOMContentLoaded", function () {
  click("skills"); // Menampilkan daftar skills saat halaman dimuat
  document.querySelector('.AboutLinkText[data-target="skills"]').classList.add("active");
});
