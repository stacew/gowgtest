const canvas = document.getElementById("gameCanvas");
document.addEventListener("contextmenu", handleCM);
function handleCM(event) {
    event.preventDefault();
}



const ctx = canvas.getContext("2d");
