import BrotherMovements from "./BrotherMovements.js";

const { ref, onMounted, inject } = Vue;


export default {
    components: {BrotherMovements},
    emits: ['changeComponent'],
    expose:['fetchAffiliations'],
    setup(props, { emit, expose })
    {
        const affiliations = inject('affiliations');
        const fetchAffiliations = inject('fetchAffiliations');
        const setAffiliation = inject('setAffiliation');
        const firstMonth = ref(1);
        const totalInstallments = ref(1);
        const fetchPeriod = () => {

            fetch('/api/periods/current')
                .then(response => {
                    if (!response.ok) {
                        if (response.status === 401) {
                            window.location.href = '/';
                        }
                    }
                    return response.json();
                })
                .then(data => {
                    firstMonth.value = data.first_month_installment;
                    totalInstallments.value = data.total_installments;
                });
        };
        const letterMonth = (month) => {
            const date = new Date(2009, month-2, 10);  // 2009-11-10
            const monthName = date.toLocaleString('default', { month: 'long' });
            return monthName[0].toUpperCase();
        };
        const notes = (affiliation) => {
            if (affiliation.honorary) {
                return "Miembro Honorario";
            }else if (affiliation.end_date != null) {
                return "Exmiembro";
            }
            return "";
        };
        onMounted(() => {
            fetchPeriod();
            fetchAffiliations();
        });
        return { affiliations, firstMonth, totalInstallments, fetchAffiliations, letterMonth, notes, fetchPeriod, setAffiliation };
        },
    methods: {
        toggle(name, affiliation) {
            this.$emit("changeComponent",name, affiliation);
        },

    },
    template: `
        <div>
            <h2>Afiliaciones</h2>
            <a @click="toggle('exaltation', null)" role="button" > <i class="bi bi-plus"></i> Exaltaci&oacute;n  </a>
            <table class="table table-bordered">
                <thead>
                    <tr class="align-items-center">
                        <th>Compañero</th>
                        <th>Notas</th>
                        <th>Deuda Vencida</th>
                        <th v-for="n in totalInstallments" >{{ letterMonth(n + firstMonth)}}</th>
                        <th>Saldo</th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="affiliation in affiliations">
                        <td>{{ affiliation.brother.first_name + " " +affiliation.brother.last_names }}</td>
                        <td>{{ notes(affiliation) }}</td>
                        <td class="text-end">{{ affiliation.overdue }}</td>
                        <td v-for="n in totalInstallments"  class="text-center">
                            <i v-if="affiliation.payedMonth.includes(n + firstMonth - 1)" class="bi bi-check-square-fill"></i>
                            <i v-else-if="affiliation.months.includes(n + firstMonth - 1)" class="bi bi-square"></i>
                        </td>
                        <td class="text-end">{{ affiliation.balance }}</td>
                        <td>
                          <a @click="setAffiliation(affiliation)" role="button" data-bs-toggle="modal" data-bs-target="#exampleModal" ><i class="bi bi-receipt"></i></a>&nbsp;
                          <a @click="toggle('brother_payment', affiliation)" role="button"><i class="bi bi-cash" ></i></a>&nbsp;
                          <a @click="toggle('Movimientos')" role="button"><i class="bi bi-cart4"></i></a>&nbsp;
                        </td>
                    </tr>
                </tbody>
            </table>
            <BrotherMovements />
        </div>`,

}

// This is the Affiliation component that is used in the Affiliation.vue file.