Vue.component('treasury-app', {
    data() {
        return {
            movements: [],
            balance: 0,
            incomes: 0,
            outcomes: 0
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
                    this.movements = []
                    let Format = new Intl.NumberFormat('es-AR',
                        { style: 'currency', currency: 'ARS', minimumFractionDigits: 2 , currencySign: 'accounting' });
                    this.balance = Format.format(data.balance);
                    this.incomes = Format.format(data.incomes);
                    this.outcomes = Format.format(data.outcomes);
                    for (let movement of data.movements) {
                        movement.date = new Date(movement.date).toLocaleDateString();
                        movement.amount =  Format.format(movement.amount);
                        this.movements.push(movement);
                    }
                });
        }
    },
    template: `
        <div>
            <h2>Tesoreria</h2>
            <h3>Resumen</h3>
            <div class="row">
                <div class="col">
                    <b>Ingresos</b>
                    <p>{{ incomes }}</p>
                </div>
                <div class="col">
                    <b>Egresos</b>
                    <p>{{ outcomes }}</p>
                </div>
                <div class="col">
                    <b>Balance</b>
                    <p>{{ balance }}</p>
                </div>
            </div>
            <table class="table">
                <thead>
                    <tr>
                        <th>Fecha</th>
                        <th>Descripci&oacute;n</th>
                        <th>Tipo</th>
                        <th>Credito</th>
                        <th>Debito</th>
                        <th>Recibo</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="movement in movements">
                        <td>{{ movement.date }}</td>
                        <td>{{ movement.description }}</td>
                        <td>{{ movement.movement_type.description }}</td>
                        <td v-if="movement.movement_type.credit">{{ movement.amount }}</td>
                        <td v-else></td>
                        <td v-if="!movement.movement_type.credit">{{ movement.amount }}</td>
                        <td v-else></td>
                        <td>{{ movement.receipt }}</td>  
                    </tr>
                </tbody></table></div>`
});

Vue.component('treasury-movement', {
    data() {
        return {
            movement: { date: '', description: '', type: '', amount: 0, receipt: '' },
            types: []
        };
    },
    mounted() {
        this.fetchTypes();
    },
    methods: {
        fetchTypes() {
            fetch('/manual-types')
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
                    this.movement = { date: '', description: '', type: '', amount: 0, receipt: '' };
                });
        }
    },
    template: `
        <div>
        <div>
            <button class="btn btn-primary" type="button" data-bs-toggle="collapse" data-bs-target="#collapseCreate" 
            aria-expanded="false" aria-controls="collapseCreate">Create Movement</button>
        </div>
        <div class="collapse" id="collapseCreate">
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
        </div>
        </div>`
});



new Vue({
    el: '#app'
});