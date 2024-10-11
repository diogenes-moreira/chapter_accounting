const { ref, inject} = Vue;


export default {
    setup(props, { emit, expose }) {
        const affiliation = inject('affiliation');
        const amount = ref('');
        const date = ref('');
        const receipt = ref('');
        const saveBrotherPayment = () => {
            fetch('/api/affiliations/payment', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    affiliation_id: affiliation.value.ID,
                    amount: amount.value,
                    date: date.value,
                    receipt: receipt.value
                })
            })
                .then(response => {
                    if (!response.ok) {
                        if (response.status === 401) {
                            window.location.href = '/';
                        }
                    }
                    return response.json();
                })
                .then(data => {
                    console.log(data);
                    toggle('affiliations');
                });
        }
        const toggle = (name) => {
            emit("changeComponent",name);
        }
        return {  amount, date, receipt, saveBrotherPayment, toggle, affiliation };
        },
    template: `
      <div>
        <h1>Brother Payment</h1>
        <form @submit.prevent="saveBrotherPayment">
          <div class="form-group">
            <p>Brother Name</p>
            <p>{{ affiliation.brother.first_name }} {{ affiliation.brother.last_names}}</p>
          </div>
          <div class="form-group">
            <label for="amount" class="form-label">Amount</label>
            <div class="input-group">
              <span class="input-group-text">$</span>
              <input type="number" min="0" step="0.01" data-number-to-fixed="2" data-number-stepfactor="100" class="form-control currency" id="amount"  v-model="amount" />
            </div>      
          </div>
          <div class="form-group" >
            <label for="date" class="form-label">Date</label>
            <input type="date" id="date" v-model="date" class="form-control date" required />
          </div>
          <div class="form-group">
            <label for="receipt" class="form-label">Recibo</label>
            <input type="text" id="receipt" v-model="receipt" class="form-control" />
          </div>
          <button type="submit" class="btn btn-primary">Save</button>
          <button type="button" @click="toggle('affiliations')" class="btn btn-secondary">Cancel</button>
        </form>
      </div>`,
}