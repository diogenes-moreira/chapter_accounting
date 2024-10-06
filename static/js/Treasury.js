Vue.component('treasury-app', {
    data() {
        return {
            movements: []
        };
    },
    mounted() {
        this.fetchTreasury();
        this.$root.$on('movement-created', data => {
            this.fetchTreasury();
        });
    },
    methods: {
        fetchTreasury() {
            // todo parametrize treasury id
            fetch('/treasury/1')
                .then(response => response.json())
                .then(data => {
                    this.movements = data.movements;
                });
        }
    },
    template: `
        <div>
            <h2>Tesoreria</h2>
            <table class="table">
                <thead>
                    <tr>
                        <th>Fecha</th>
                        <th>Descripci&oacute;n</th>
                        <th>Tipo</th>
                        <th>Credito</th>
                        <th>Debito</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="movement in movements">
                        <td>{{ movement.date }}</td>
                        <td>{{ movement.description }}</td>
                        <td>{{ movement.type }}</td>
                        <td v-if="movement.credit">{{ movement.amount }}</td>
                        <td v-else></td>
                        <td v-if="!movement.debit">{{ movement.amount }}</td>
                        <td v-else></td>  
                    </tr>
                </tbody></table></div>`
});

Vue.component('treasury-movement', {
    data() {
        return {
            movement: { date: '', description: '', type: '', amount: 0 },
            types: []
        };
    },
    mounted() {
        this.fetchTypes();
    },
    methods: {
        fetchTypes() {
            fetch('/movement-types')
                .then(response => response.json())
                .then(data => {
                    this.types = data;
                });
        },
        createMovement() {
            fetch('/treasury/1/movements', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(this.movement)
            })
                .then(response => response.json())
                .then(data => {
                    this.$root.$emit('movement-created', data);
                    this.movement = { date: '', description: '', type: '', amount: 0 };
                });
        }
    },
    template: `
        <div>
            <h2>Create Movement</h2>
            <form @submit.prevent="createMovement">
                <div class="form-group">
                    <label for="date">Date</label>
                    <input type="date" v-model="movement.date" class="form-control" id="date" required>
                </div>
                <div class="form-group">
                    <label for="description">Description</label>
                    <input type="text" v-model="movement.description" class="form-control" id="description" required>
                </div>
                <div class="form-group">
                    <label for="type">Type</label>
                    <select v-model="movement.type" class="form-control" id="type" required>
                        <option v-for="type in types" :key="type.ID" :value="type.description">{{ type.description }}</option>
                    </select>
                </div>
                <div class="form-group">
                    <label for="amount">Amount</label>
                    <input type="number" v-model="movement.amount" class="form-control" id="amount" required>
                </div>
                <button type="submit" class="btn btn-primary">Create</button>
            </form>
        </div>`
});



new Vue({
    el: '#app'
});