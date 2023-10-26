import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        isLogged: false,
    },
    mutations: {
        setLogged(state, payload) {
            state.isLogged = payload
        }
    },
})