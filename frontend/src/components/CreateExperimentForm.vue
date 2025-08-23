<template>
    <div class="form-container">
        <h2>Create New Experiment</h2>
        <form @submit.prevent="handleSubmit">
            <div class="form-group">
                <label for="name">Experiment Name</label>
                <input type="text" id="name" v-model="experimentName" required />
            </div>
        
            <div class="form-group">
                <label>Variations (at least 2)</label>
                <div v-for="(variation, index) in variations" :key="index" class="variation-input">
                    <input type="text" v-model="variations[index]" required />
                    <button type="button" @click="removeVariation(index)" v-if="variations.length > 2">Remove</button>
                </div>
                <button type="button" @click="addVariation">Add Variation</button>
            </div>

            <button type="submit" :disabled="store.isLoading">
                {{ store.isLoading ? 'Creating...' : 'Create Experiment'}}
            </button>
            <p v-if="errorMessage" class="error"> {{ errorMessage }} </p>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { useExperimentStore } from '@/stores/experimentStore'

const store = useExperimentStore();

const experimentName= ref('');
// start with two empty variation fields.
const variations = ref(['', '']);
const errorMessage = ref('');

const addVariation = () => {
    variations.value.push('');
};

const removeVariation = (index) => {
    variations.value.splice(index, 1);
};

const handleSubmit = async () => {
    errorMessage.value = '';
    const experimentData = {
        name: experimentName.value,
        // filter out any empty variation fields before submitting.
        variations: variations.value.filter(v => v.trim() !== ''),
    };

    if (experimentData.variations.length < 2) {
        errorMessage.value = 'Please provide at least two variations.';
        return;
    }



    try {
        await store.addExperiment(experimentData);
        //Reset the form on seccesfull submission.
        experimentName.value = '';
        variations.value = ['', ''];
    } catch (error) {
        errorMessage.value = 'An error occurred. Please try again.';
    }
};
</script>

<style scoped>
.form-group {
    margin-bottom: 20px;
}
label {
    display: block;
    margin-bottom: 5px;
}
input[type="text"] {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
}
.variation-input {
    display: flex;
    gap: 10px;
    margin-bottom: 10px;
}
.error {
    color: red;
    margin-top: 10px;
}
</style>