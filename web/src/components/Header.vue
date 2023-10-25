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
            const user_id = '6536bd96345f7de7d8b9604c'
            fetch(`${config.API_URL}/users/${user_id}`)
                .then(response => response.json())
                .then(data => {
                    this.founds = data.founds
                })
        },
        endDay() {
            fetch(`${config.API_URL}/stocks/random`)
                .then(response => response.json())
                .then(() => {
                    alert('Novo dia iniciado com sucesso!')
                    
                    
                })
                .catch(error => {
                    alert('ERROR ao iniciar um novo dia', error)
                })
        }
    },
    created() {
        this.getUserData()
    }
}
</script>

<style>

</style>