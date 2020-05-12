const canvas = document.getElementById("jsCanvas");
const ctx = canvas.getContext("2d");
const colors = document.getElementsByClassName("jsColor");
const range = document.getElementById("jsRange");
const modeBtn = document.getElementById("jsMode");
const saveBtn = document.getElementById("jsSave");

const CANVAS_SIZE = 700;

//css로 정해준 canvas pixel 사이즈를 알아야 그릴 때 위치가 맞음.
canvas.width = CANVAS_SIZE;
canvas.height = CANVAS_SIZE;

ctx.fillStyle = "white";
ctx.fillRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);//캔버스 저장 시 투명 배경 제거


const INITIAL_COLOR = "#2c2c2c";
ctx.strokeStyle = INITIAL_COLOR;
ctx.fillStyle = INITIAL_COLOR;
ctx.lineWidth = 6;

let painting = false;
let fillMode = false;

function startPainting() {
    painting = true;
}
function stopPainting() {
    painting = false;
}

function onMouseMove(event) {
    const x = event.offsetX;
    const y = event.offsetY;
    if (!painting) {
        ctx.beginPath();
        ctx.moveTo(x, y);
    } else if (fillMode === false) {
        ctx.lineTo(x, y);
        ctx.stroke(); //그리기 : 선을 긋고. Line 경로 종료?
    }
}
// ctx.beginPath(); //Line To 기록 초기화
// ctx.closePath(); //올가미 도구 느낌. 기록이 시작된 경로의 첫점으로 닫는 느낌.


function handleColorClick(event) {
    color = event.target.style.backgroundColor;
    ctx.strokeStyle = color;
    ctx.fillStyle = color;
}

function handleRangeChange(event) {
    ctx.lineWidth = event.target.value;
}

function handleModeClick(event) {
    if (fillMode === true) {
        fillMode = false;
        modeBtn.innerText = "Paint";
    } else {
        fillMode = true;
        modeBtn.innerText = "Fill";
    }
}

function handleCanvasClick(event) {
    if (fillMode)
        ctx.fillRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);
}

function handleSaveClick(event) {
    const image = canvas.toDataURL();//default : png
    // const image = canvas.toDataURL("image/jpeg");
    const link = document.createElement("a");
    link.href = image;
    link.download = "paintJS[EXPORT]";
    link.click();
}

function handleCM(event) {
    event.preventDefault(); //우측 마우스 방지
}


//addEventListener
if (canvas) {
    canvas.addEventListener("mousemove", onMouseMove);
    canvas.addEventListener("mousedown", startPainting);
    canvas.addEventListener("mouseup", stopPainting);
    canvas.addEventListener("mouseleave", stopPainting);
    canvas.addEventListener("click", handleCanvasClick);
    canvas.addEventListener("contextmenu", handleCM);
}
Array.from(colors).forEach(colorIndex => colorIndex.addEventListener("click", handleColorClick))
if (range) {
    range.addEventListener("input", handleRangeChange);
}
if (modeBtn) {
    modeBtn.addEventListener("click", handleModeClick);
}
if (saveBtn) {
    saveBtn.addEventListener("click", handleSaveClick);
}