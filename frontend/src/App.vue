<template>
  <div id="app-container">
    <header>
      <h1>A/B Testing Dashboard</h1>
    </header>
    <main>
      <CreateExperimentForm />
      <ExperimentList />
    </main>
  </div>
</template>

<script setup>
import ExperimentList from './components/ExperimentList.vue';
import CreateExperimentForm from './components/CreateExperimentForm.vue';
</script>

<style>
#app-container {
  max-width: 960px;
  margin: 0 auto;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  color: #2c3e50;
}
header {
  border-bottom: 2px solid #eee;
  padding-bottom: 20px;
  margin-bottom: 30px;
}
main {
  display: flex;
  gap: 40px;
}
main > * {
  flex: 1;
}
</style>