function drawPlayer() {
    ctx.beginPath();
    ctx.arc(playerX, playerY, ballRadius, 0, Math.PI * 2);
    ctx.fillStyle = "#0095DD";
    ctx.fill();
    ctx.closePath();
}

function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    drawPlayer();
    requestAnimationFrame(draw);// Call Draw Only Server
}
draw();