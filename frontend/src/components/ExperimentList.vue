<template>
  <div class="experiment-list">
    <h2>Current Experiments</h2>
    <div v-if="store.isLoading">Loading...</div>
    <div v-else-if="store.error" class="error">{{ store.error }}</div>
    <ul v-else-if="store.experiments.length > 0" class="experiment-cards">
      <li v-for="exp in store.experiments" :key="exp.ID" class="experiment-card">
        <h3>{{ exp.Name }}</h3>
        
        <button @click="simulateUserJourney(exp.ID)" class="convert-button">
          Run User
        </button>
        <!-- Table to show variation results -->
        <table>
          <thead>
            <tr>
              <th>Variation</th>
              <th>Participants</th>
              <th>Conversions</th>
              <th>Rate</th>
              <th>Simulate</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="variation in exp.Variations" :key="variation.ID">
              <td>{{ variation.Name }}</td>
              <td>{{ variation.Participants }}</td>
              <td>{{ variation.Conversions }}</td>
              <td>{{ calculateConversionRate(variation) }}%</td>
            </tr>
          </tbody>
        </table>
      </li>
    </ul>
    <div v-else>No experiments found. Create one to get started!</div>
  </div>
</template>


<script setup>
import { onMounted } from 'vue';
import { useExperimentStore } from '@/stores/experimentStore';
import api from '@/services/api';

const store = useExperimentStore();

onMounted(() => {
  store.fetchExperiments();
});

const calculateConversionRate = (variation) => {
  if (variation.Participants === 0) {
    return '0.00'; // Return a string to be consistent with toFixed()
  }
  const rate = (variation.Conversions / variation.Participants) * 100;
  return rate.toFixed(2);
};

// --- REWRITTEN LOGIC ---
// This function simulates a full user journey.
const simulateUserJourney = async (experimentId) => {
  try {
    // Step 1: Assign a user to a variation for this experiment.
    // The backend will randomly pick one and increment its participant count.
    const assignmentResponse = await api.assignToVariation(experimentId);
    const assignedVariationId = assignmentResponse.data.variationId;

    await api.recordConversion(assignedVariationId);

    // Step 3: Refresh the data from the server to show the updated numbers.
    await store.fetchExperiments();

  } catch (error) {
    console.error("Failed to simulate user journey:", error);
  }
};
</script>

<style scoped>
/* Scoped styles apply only to this component */
.experiment-list {
  border-left: 2px solid #eee;
  padding-left: 40px;
}
.experiment-cards {
  list-style-type: none;
  padding: 0;
}
.experiment-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 15px;
}
th, td {
  text-align: left;
  padding: 8px;
  border-bottom: 1px solid #ddd;
}
th {
  background-color: #f9f9f9;
}
.convert-button {
  padding: 5px 10px;
  font-size: 12px;
  cursor: pointer;
  border: 1px solid #007bff;
  background-color: #fff;
  color: #007bff;
  border-radius: 4px;
}
.convert-button:hover {
  background-color: #007bff;
  color: #fff;
}
.error {
  color: red;
}
</style>