<script setup lang="ts">
import { ref } from "vue";

const inputs = ref([{ value: "" }]); // array of input fields
const results = ref([]);

const addInput = () => {
  inputs.value.push({ value: "" });
};

const calculate = async () => {
  try {
    // preparing data to send to the API
    const data = inputs.value.map((input) => ({ input: input.value }));

    // call the API
    const response = await fetch("http://localhost:8080/process", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });

    // throw error if response is not ok
    if (!response.ok) {
      throw new Error("Error: " + (await response.text()));
    }

    const result = await response.json();
    results.value = result.results;
  } catch (error: any) {
    console.error("Error:", error.message);
    alert(error.message);
  }
};

const handleSubmit = (event: Event) => {
  event.preventDefault();
  const form = event.target as HTMLFormElement;

  if (form.checkValidity()) {
    calculate();
  } else {
    form.reportValidity();
  }
};
</script>

<template>
  <div class="container">
    <h1>Salary Calculator</h1>

    <div class="form-wrapper">
      <form @submit.prevent="handleSubmit" class="salary-form">
        <!-- input fields -->
        <div v-for="(input, index) in inputs" :key="index" class="input-group">
          <input
            v-model="input.value"
            type="text"
            placeholder="Enter input"
            class="input-text"
            required
          />
          <div v-if="results.length" class="result-display">
            <span>{{ results[index] }}</span>
          </div>
        </div>

        <!-- buttons to add input and calculate -->
        <div class="button-group">
          <button type="button" @click="addInput" class="button-add">
            Add Input
          </button>
          <button type="submit" class="button-calculate">Calculate</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.salary-form {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.input-group {
  display: flex;
  align-items: center;
  margin-bottom: 1rem;
}

.input-text {
  padding: 0.5em;
  font-size: 1em;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.result-display {
  margin-left: 1em;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.button-group {
  margin-top: 1rem;
  display: flex;
  gap: 1rem;
}

button {
  border-radius: 8px;
  border: 1px solid transparent;
  padding: 0.6em 1.2em;
  font-size: 1em;
  font-weight: 500;
  font-family: inherit;
  background-color: #6a1b9a;
  color: #ffffff;
  cursor: pointer;
  transition: border-color 0.25s, background-color 0.25s;
}

button:hover {
  background-color: #4a0072;
}
</style>
