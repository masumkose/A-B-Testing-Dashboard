<template>
    <div class="experiment-list">
        <h2>Current Experiments</h2>
        <div v-if="store.isLoading">Loading...</div>
        <div v-else-if="store.error" class="error">{{ store.error }}</div>
        <ul v-else-if="store.experiments.length > 0">
            <li v-for="exp in store.experiments" :key="exp.ID">
                <strong> {{exp.Name}}</strong>
                <ul>
                    <li v-for="variation in exp.Variations" :key="variation.ID">
                        {{variation.Name}}
                    </li>
                </ul>
            </li>
        </ul>
        <div v-else>No experiment found. Create one to get started!</div>
    </div>
</template>


<script setup>
import { onMounted } from 'vue';
import { useExperimentStore } from '@/stores/experimentStore';

const store = useExperimentStore();

onMounted(() => {
    store.fetchExperiments();
});
</script>

<style scoped>
.experiment-list {
    border-left: 2px solid #eee;
    padding-left: 40px;
}
ul {
    list-style-type: none;
    padding: 0;
}
li {
    margin-bottom: 15px;
}
.error {
    color: red;
}
</style>