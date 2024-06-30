function toggleNav() {
    const btn = document.querySelector(".navmob-btn");
    const nav = document.querySelector(".navmob");
    const content = document.querySelector(".content-container");

    btn.classList.toggle('open');
    nav.classList.toggle('open');
    content.classList.toggle('open');
}