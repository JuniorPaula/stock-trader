<template>
    <v-toolbar app>
        <v-toolbar-title class="headline text-upercase mr-4">
            <span>Stock</span>
            <span class="font-weight-light">Trader</span>
        </v-toolbar-title>
        <v-toolbar-items>
            <v-btn flat to="/">Inicio</v-btn>
            <v-btn flat to="/stocks">Ações</v-btn>
            <v-btn flat to="/portfolio">Portfolio</v-btn>
        </v-toolbar-items>

        <v-spacer></v-spacer>

        <v-toolbar-items>
            <v-btn @click="endDay" flat>Finalizar Dia</v-btn>
            <v-layout align-center>
                <span class="text-uppercase grey--text text-darken-2">
                    Saldo: {{ founds | currency  }}
                </span>
            </v-layout>
            <v-btn @click="logout" flat>Sair</v-btn>
        </v-toolbar-items>
    </v-toolbar>
</template>

<script>
import config from '@/config/config';

export default {
    data() {
        return {
            founds: 0
        }
    },
    methods: {
        getUserData() {
            const useData = localStorage.getItem('__user__')
            const user = JSON.parse(useData)

            fetch(`${config.API_URL}/api/users/${user.user_id}`, {
                headers: {
                    'Authorization': `Bearer ${user.token}`
                }
            })
            .then(response => response.json())
            .then(data => {
                this.founds = data.founds
            })
        },
        endDay() {
            const useData = localStorage.getItem('__user__')
            const user = JSON.parse(useData)

            fetch(`${config.API_URL}/api/stocks/random`, {
                headers: {
                    'Authorization': `Bearer ${user.token}`
                }
            })
            .then(response => response.json())
            .then(() => {
                alert('Novo dia iniciado com sucesso!')
                
                
            })
            .catch(error => {
                alert('ERROR ao iniciar um novo dia', error)
            })
        },
        logout() {
            this.$store.commit('setLogged', false)
            localStorage.removeItem('__user__')
            this.$router.push('/login')
        }
    },
    created() {
        this.getUserData()
    }
}
</script>

<style>

</style>