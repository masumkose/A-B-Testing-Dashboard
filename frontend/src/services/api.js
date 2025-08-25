import axios from "axios";

const apiClient = axios.create({
    baseURL: 'http://localhost:8080/api',
    headers: {
        'Content-Type' : 'application/json'
    }
});

export default {
    getExperiments() {
        return apiClient.get('/experiments');
    },
    createExperiments(experimentData) {
        return apiClient.post('/experiments', experimentData);
    },
    recordConversion(variationID) {
        // This endpoint takes the variation ID in the URL.
        return apiClient.post(`/variations/${variationID}/convert`);
    },
    assignToVariation(experimentID) {
        // This endpoint return the variation the use was assigned to.
        return apiClient.post(`/experiments/${experimentID}/assign`);
    }
};