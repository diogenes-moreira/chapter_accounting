Vue.component('affiliation-app', {
    data() {
        return {
            affiliations: [],
            firstMonth: 1
        };
    },
    mounted() {
        this.fetchAffiliations();
    },
    methods: {
        fetchAffiliations() {
            fetch('/chapters/1/affiliations/1')
                .then(response => response.json())
                .then(data => {
                    this.affiliations = data;
                    let Format = new Intl.NumberFormat('es-AR',
                        { style: 'currency', currency: 'ARS', minimumFractionDigits: 2 , currencySign: 'accounting' });
                    this.balance = Format.format(data.balance);
                    for (let affiliation of this.affiliations) {
                        affiliation.overdue = Format.format(affiliation.overdue);
                        affiliation.balance = Format.format(affiliation.balance);
                        affiliation.payedMonth = [];
                        try {
                            for (let installment of affiliation.installments) {
                                if (installment.paid) {
                                    if (installment.month != 0) {
                                        affiliation.payedMonth.push(installment.month);
                                    }
                                }
                            }
                        }catch (e) {
                            console.log(e);
                        }
                    }
                });
            fetch('/api/periods/1')
                .then(response => response.json())
                .then(data => {
                    this.firstMonth = data.first_month_installment;
                });
        },
        letterMonth(month) {
            const date = new Date(2009, month-1, 10);  // 2009-11-10
            const monthName = date.toLocaleString('default', { month: 'long' });
            return monthName[0].toUpperCase();
        },
        notes(affiliation) {
            if (affiliation.honorary) {
                return "Miembro Honorario";
            }else if (affiliation.end_date != null) {
                return "Exmiembro";
            }
            return "";

        }

    },
    template: `
        <div>
            <h2>Afiliaciones</h2>
            <table class="table table-bordered">
                <thead>
                    <tr class="align-items-center">
                        <th>Compañero</th>
                        <th>Notas</th>
                        <th>Deuda Vencida</th>
                        <th v-for="n in 12" v-if="n >= firstMonth">{{ letterMonth(n)}}</th>
                        <th>Saldo</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="affiliation in affiliations">
                        <td>{{ affiliation.brother.first_name + " " +affiliation.brother.last_names }}</td>
                        <td>{{ notes(affiliation) }}</td>
                        <td class="text-end">{{ affiliation.overdue }}</td>
                        <td v-for="n in 12" v-if="n >= firstMonth" class="text-center">
                            <i v-if="affiliation.payedMonth.includes(n)" class="bi bi-check-square-fill"></i>
                            <i v-else class="bi bi-square"></i>
                        </td>
                        <td class="text-end">{{ affiliation.balance }}</td>
                    </tr>
                </tbody>
            </table>
        </div>`
});

new Vue({
    el: '#app',
    data: {
        currentComponent: 'affiliation-app'
    }
});