<template>
    <div>
        <h1 class="display-3 font-weight-light mb-4">Negocie e Consulte suas Ações</h1>
        <v-sheet :elevation="6" class="pa-2" color="primary">
            <v-icon class="white--text mr-3">info</v-icon>
            <span class="headline white--text font-weight-light">Você pode Salvar & Carregar os Dados</span>
        </v-sheet>
        <v-sheet :elevation="6" class="pa-2 mt-3" color="success darken-1">
            <v-icon class="white--text mr-3">info</v-icon>
            <span class="headline white--text font-weight-light">Click em 'Finalizar Dia' para iniciar um novo dia.</span>
        </v-sheet>
        <v-divider class="my-4" />
        <p class="display-1"><strong>Seu Saldo:</strong> {{ founds }}</p>
    </div>
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
        }
    },
    created() {
        this.getUserData()
    }
}
</script>

<style></style>