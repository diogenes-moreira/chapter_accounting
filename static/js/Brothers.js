Vue.component('brothers-app', {
    data() {
        return {
            brothers: []
        };
    },
    mounted() {
        this.fetchBrothers();
    },
    methods: {
        fetchBrothers() {
            fetch('/brothers')
                .then(response => response.json())
                .then(data => {
                    this.brothers = data;
                });
        }
    },
    template: `
        <div>
            <h2>Hermanos</h2>
            <table class="table">
                <thead>
                    <tr>
                        <th>Nombre</th>
                        <th>Apellido</th>
                        <th>Documento</th>
                        <th>Fecha de Nacimiento</th>
                        <th>Fecha de Ingreso</th>
                        <th>Fecha de Egreso</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="brother in brothers">
                        <td>{{ brother.name }}</td>
                        <td>{{ brother.last_name }}</td>
                        <td>{{ brother.document }}</td>
                        <td>{{ brother.birth_date }}</td>
                        <td>{{ brother.entry_date }}</td>
                        <td>{{ brother.exit_date }}</td>
                    </tr>
                </tbody>
            </table>
        </div>`

});

new Vue({
    el: '#app'
});