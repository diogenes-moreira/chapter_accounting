const { shallowRef, provide, ref } = Vue;
import Affiliations from '/js/Affiliation.js';
import Exaltation from "/js/Exaltation.js";
import BrotherPayment from "/js/BrotherPayment.js";

export default {
        name: 'root-component',
        setup() {
            const current = shallowRef(Affiliations);
            const affiliation = ref({movements:[]});
            const affiliations = ref([]);
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
                        for (let affiliation of affiliations.value) {
                            affiliation.overdue = Format.format(affiliation.overdue);
                            affiliation.balance = Format.format(affiliation.balance);
                            affiliation.payedMonth = [];
                            affiliation.months = [];
                            try {
                                for (let installment of affiliation.installments) {
                                    if (installment.month !== 0) {
                                        affiliation.months.push(installment.month);
                                        if (installment.paid) {
                                            affiliation.payedMonth.push(installment.month);
                                        }
                                    }
                                }
                            }catch (e) {
                                console.log(e);
                            }
                        }
                    });
            }
            const setAffiliation = (a) => {
                affiliation.value = a;
            }
            provide('fetchAffiliations', fetchAffiliations);
            provide('affiliations', affiliations);
            provide('affiliation', affiliation);
            provide('setAffiliation',setAffiliation );
            return { current, affiliation, affiliations, fetchAffiliations, setAffiliation };
        },
        components: {
            Affiliations,
            Exaltation,
            BrotherPayment,
        },
        methods: {
            changeComponent(name, affiliation) {
                if (name === 'affiliations') {
                    this.current = Affiliations;
                    this.fetchAffiliations();
                } else if (name === 'exaltation') {
                    this.current = Exaltation;
                } else if (name === 'brother_payment') {
                    this.current = BrotherPayment;
                    this.affiliation = affiliation;
            }
        }}
    ,
        template: `
              <div class="demo">
                <KeepAlive>
                  <component :is="current" 
                             @change-component="(name, affiliation)=>{changeComponent(name,affiliation)}">
                  </component>
                </KeepAlive>
              </div>
            `
}

