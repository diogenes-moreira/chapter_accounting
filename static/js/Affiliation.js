const { ref, onMounted } = Vue;

export default {
    setup() {
        const affiliations = ref([]);
        const firstMonth = ref(1);
        const totalInstallments = ref(1);
        const fetchAffiliations = () => {
            fetch('/api/chapters/affiliations')
                .then(response => {
                    if (!response.ok) {
                        if (response.status === 401) {
                            window.location.href = '/';
                        }
                    }
                    return response.json();
                })
                .then(data => {
                    affiliations.value = data;
                    let Format = new Intl.NumberFormat('es-AR',
                        { style: 'currency', currency: 'ARS', minimumFractionDigits: 2 , currencySign: 'accounting' });
                    affiliations.balance = Format.format(data.balance);
                    for (let affiliation of affiliations.value) {
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
            fetchAffiliations();
        });
        return { affiliations, firstMonth, totalInstallments, fetchAffiliations, letterMonth, notes };
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
                        <th v-for="n in totalInstallments" >{{ letterMonth(n + firstMonth)}}</th>
                        <th>Saldo</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="affiliation in affiliations">
                        <td>{{ affiliation.brother.first_name + " " +affiliation.brother.last_names }}</td>
                        <td>{{ notes(affiliation) }}</td>
                        <td class="text-end">{{ affiliation.overdue }}</td>
                        <td v-for="n in totalInstallments"  class="text-center">
                            <i v-if="affiliation.payedMonth.includes(n + firstMonth - 1)" class="bi bi-check-square-fill"></i>
                            <i v-else class="bi bi-square"></i>
                        </td>
                        <td class="text-end">{{ affiliation.balance }}</td>
                    </tr>
                </tbody>
            </table>
        </div>`,

}

// This is the Affiliation component that is used in the Affiliation.vue file.