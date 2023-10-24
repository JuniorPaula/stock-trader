import config from '../../config/config'

export default {
    state: {
        stocks: []
    },
    mutations: {
        setStocks(state, stocks) {
            state.stocks = stocks;
        },

    },
    actions: {
        buyStock({ commit }, order) {
            console.log(order)
            commit();
        },
        initStocks({ commit }) {
            fetch(`${config.API_URL}/stocks`)
                .then(response => response.json())
                .then(data => {
                    commit('setStocks', data);
                })
        },
    },
    getters: {
        stocks(state) {
            return state.stocks;
        },
    }
}