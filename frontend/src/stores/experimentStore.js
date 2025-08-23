import { defineStore } from "pinia";
import api from '@/services/api';

// 'use experimentstore' is the conventional name for the hook.
export const useExperimentStore = defineStore('experiments', {
    // state : The data our store holds.
    state: () => ({
        experiments: [],
        isLoading: false,
        error: null,
    }),

    // actions: Functions that can change the state They are often asynchronous.
    actions: {
        async fetchExperiments() {
            this.isLoading = true;
            this.error = null;
            try {
                const response = await api.getExperiments();
                this.experiments = response.data; // update state with data from the API
            } catch (error) {
                this.error = 'Failed to fetch experiments.';
                console.error(error);
            } finally {
                this.isLoading = false;
            }
        },

        async addExperiment(experimentData) {
            this.isLoading = true;
            this.error = null;
            try {
                const response = await api.createExperiments(experimentData);
                // add the newly created experiment to our local state without needing to re-fetch everything.
                this.experiments.push(response.data);
            } catch (error) {
                this.error = 'Failed to create experiment.';
                console.error(error); // re throw the error so the component knows it failed.
                throw error;
            } finally {
                this.isLoading = false;
            }
        },
    },
});