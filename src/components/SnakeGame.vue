<template>
  <div class="game-container">
    <canvas ref="gameCanvas" width="400" height="400"></canvas>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';

const gameCanvas = ref<HTMLCanvasElement | null>(null);

interface Position {
  x: number;
  y: number;
}

interface Snake {
  x: number;
  y: number;
  dx: number;
  dy: number;
  cells: Position[];
  maxCells: number;
}

const grid = 16;
let count = 0;
let animationId: number | null = null;

const snake: Snake = {
  x: 160,
  y: 160,
  dx: grid,
  dy: 0,
  cells: [],
  maxCells: 4
};

const apple: Position = {
  x: 320,
  y: 320
};

function getRandomInt(min: number, max: number): number {
  return Math.floor(Math.random() * (max - min)) + min;
}

function loop() {
  if (!gameCanvas.value) return;
  
  animationId = requestAnimationFrame(loop);
  const context = gameCanvas.value.getContext('2d');
  if (!context) return;

  // slow game loop to 15 fps instead of 60 (60/15 = 4)
  if (++count < 4) {
    return;
  }

  count = 0;
  context.clearRect(0, 0, gameCanvas.value.width, gameCanvas.value.height);

  // move snake by its velocity
  snake.x += snake.dx;
  snake.y += snake.dy;

  // wrap snake position horizontally on edge of screen
  if (snake.x < 0) {
    snake.x = gameCanvas.value.width - grid;
  } else if (snake.x >= gameCanvas.value.width) {
    snake.x = 0;
  }

  // wrap snake position vertically on edge of screen
  if (snake.y < 0) {
    snake.y = gameCanvas.value.height - grid;
  } else if (snake.y >= gameCanvas.value.height) {
    snake.y = 0;
  }

  // keep track of where snake has been. front of the array is always the head
  snake.cells.unshift({ x: snake.x, y: snake.y });

  // remove cells as we move away from them
  if (snake.cells.length > snake.maxCells) {
    snake.cells.pop();
  }

  // draw apple
  context.fillStyle = 'red';
  context.fillRect(apple.x, apple.y, grid - 1, grid - 1);

  // draw snake one cell at a time
  context.fillStyle = 'white';
  snake.cells.forEach((cell, index) => {
    // drawing 1 px smaller than the grid creates a grid effect in the snake body
    context.fillRect(cell.x, cell.y, grid - 1, grid - 1);

    // snake ate apple
    if (cell.x === apple.x && cell.y === apple.y) {
      snake.maxCells++;

      // canvas is 400x400 which is 25x25 grids
      apple.x = getRandomInt(0, 25) * grid;
      apple.y = getRandomInt(0, 25) * grid;
    }

    // check collision with all cells after this one (modified bubble sort)
    for (let i = index + 1; i < snake.cells.length; i++) {
      // snake occupies same space as a body part. reset game
      if (cell.x === snake.cells[i].x && cell.y === snake.cells[i].y) {
        snake.x = 160;
        snake.y = 160;
        snake.cells = [];
        snake.maxCells = 4;
        snake.dx = grid;
        snake.dy = 0;

        apple.x = getRandomInt(0, 25) * grid;
        apple.y = getRandomInt(0, 25) * grid;
      }
    }
  });
}

function handleKeyDown(e: KeyboardEvent) {
  // prevent snake from backtracking on itself by checking that it's
  // not already moving on the same axis

  // left arrow key
  if (e.which === 37 && snake.dx === 0) {
    snake.dx = -grid;
    snake.dy = 0;
  }
  // up arrow key
  else if (e.which === 38 && snake.dy === 0) {
    snake.dy = -grid;
    snake.dx = 0;
  }
  // right arrow key
  else if (e.which === 39 && snake.dx === 0) {
    snake.dx = grid;
    snake.dy = 0;
  }
  // down arrow key
  else if (e.which === 40 && snake.dy === 0) {
    snake.dy = grid;
    snake.dx = 0;
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleKeyDown);
  animationId = requestAnimationFrame(loop);
});

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeyDown);
  if (animationId !== null) {
    cancelAnimationFrame(animationId);
  }
});
</script>

<style scoped>
.game-container {
  /* height: 100vh; */
  margin: 0;
  background: rgba(77, 72, 160);
  display: flex;
  align-items: center;
  justify-content: center;
  border: 4px solid black;
}

canvas {
  border: 1px solid white;
}
</style>