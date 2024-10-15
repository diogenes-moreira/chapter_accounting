import CompanionMovements from "./CompanionMovements.js";

const { ref, onMounted, inject } = Vue;


export default {
    components: {CompanionMovements},
    emits: ['changeComponent'],
    expose:['fetchAffiliations'],
    setup(props, { emit, expose })
    {
        const affiliations = inject('affiliations');
        const fetchAffiliations = inject('fetchAffiliations');
        const setAffiliation = inject('setAffiliation');
        const firstMonth = ref(1);
        const totalInstallments = ref(1);
        const deposits = ref([]);
        const showDetail = (movement) => {
            deposits.value = [];
            let Format = new Intl.NumberFormat('es-AR',
                { style: 'currency', currency: 'ARS', minimumFractionDigits: 2 , currencySign: 'accounting' });
            let installments = movement.installments.filter(installment => installment.deposit != null);
            for ( let installment of installments){
                let deposit = {};
                deposit.date = new Date(installment.deposit.deposit_date).toLocaleDateString();
                deposit.amount = Format.format(installment.deposit.amount);
                deposit.description = installment.description;
                deposits.value.push(deposit);
            }

        }
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
        return { affiliations,
            firstMonth,
            totalInstallments,
            deposits,
            fetchAffiliations,
            letterMonth,
            notes,
            fetchPeriod,
            setAffiliation,
            showDetail };
        },
    methods: {
        toggle(name, param ) {
            this.$emit("changeComponent",name, param);
        },
        readOnly() {
            return window.readOnly;
        }

    },
    template: `
        <div>
            <h2>Afiliaciones</h2>
            <a v-if="!readOnly()" @click="toggle('exaltation', true)" role="button" ><i class="bi bi-plus"></i>Exaltaci&oacute;n</a>&nbsp;&nbsp;
          <a v-if="!readOnly()" @click="toggle('exaltation', false)" role="button" ><i class="bi bi-plus"></i>Afiliaci&oacute;n</a>
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
                        <td>{{ affiliation.companion.first_name + " " +affiliation.companion.last_names }}</td>
                        <td>{{ notes(affiliation) }}</td>
                        <td class="text-end">{{ affiliation.overdue }}</td>
                        <td v-for="n in totalInstallments"  class="text-center">
                            <i v-if="affiliation.payedMonth.includes(n + firstMonth - 1)" class="bi bi-check-square-fill"></i>
                            <i v-else-if="affiliation.months.includes(n + firstMonth - 1)" class="bi bi-square"></i>
                        </td>
                        <td class="text-end">{{ affiliation.balance }}</td>
                        <td>
                          <a @click="setAffiliation(affiliation)" role="button" data-bs-toggle="modal" data-bs-target="#exampleModal" ><i class="bi bi-receipt"></i></a>&nbsp;
                          <a @click="showDetail(affiliation)" role="button" data-bs-toggle="modal" data-bs-target="#detailDeposit" ><i class="bi bi-bank"></i></a>&nbsp; 
                          <a @click="toggle('companion_payment', affiliation)" role="button"><i class="bi bi-cash" ></i></a>&nbsp;
                          <a @click="toggle('companion_expenses', affiliation)" role="button"><i class="bi bi-cart4"></i></a>&nbsp;
                        </td>
                    </tr>
                </tbody>
            </table>
            <CompanionMovements />
            <div class="modal fade" id="detailDeposit" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
                <div class="modal-dialog modal-dialog-scrollable">
                    <div class="modal-content">
                    <div class="modal-header">
                        <p class="modal-title" id="exampleModalLabel">Depositos</p>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                        <table class="table">
                            <thead>
                                <tr>
                                    <th>Fecha</th>
                                    <th>Recibo</th>
                                    <th>Monto</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="movement in deposits">
                                    <td>{{ movement.date }}</td>
                                    <td>{{ movement.description }}</td>
                                    <td>{{ movement.amount }}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                      <div class="modal-footer">
                            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cerrar</button>
                      </div>
                    </div>
                </div>
              </div>
        </div>`,

}

