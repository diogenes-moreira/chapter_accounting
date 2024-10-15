const { ref, provide, onMounted } = Vue;
import ManualMovements from "./ManualMovements.js";

export default {
    setup() {
        const movements = ref([]);
        const balance = ref(0);
        const incomes = ref(0);
        const outcomes = ref(0);
        const fetchTreasury = () => {
            fetch('/api/chapters/treasury')
                .then(response => {
                    if (!response.ok) {
                        if (response.status === 401) {
                            window.location.href = '/';
                        }
                    }
                    return response.json();
                })
                .then(data => {
                    movements.value = []
                    let Format = new Intl.NumberFormat('es-AR',
                        { style: 'currency', currency: 'ARS', minimumFractionDigits: 2 , currencySign: 'accounting' });
                    balance.value = Format.format(data.balance);
                    incomes.value = Format.format(data.incomes);
                    outcomes.value = Format.format(data.outcomes);
                    for (let movement of data.movements) {
                        movement.date = new Date(movement.date).toLocaleDateString();
                        movement.amount =  Format.format(movement.amount);
                        movements.value.push(movement);
                    }
                });
        };
        onMounted(() => {
            fetchTreasury();
        });
        provide('fetchTreasury', fetchTreasury);

        return { movements, balance, incomes, outcomes, fetchTreasury };
    },
    components: { ManualMovements },
    methods: {
        readOnly:  () => {
            return window.readOnly;
        }
    },
    template: `<div>
            <h2>Tesoreria</h2>
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
            <a v-if="!readOnly()" role="button" data-bs-toggle="modal" data-bs-target="#exampleModal"  ><i class="bi bi-plus"></i>Movimiento</a>&nbsp;&nbsp;
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
                </tbody></table></div>
    <ManualMovements  @refresh="fetchTreasury"/>`
};
