const { ref, inject, onMounted} = Vue;


export default {
    setup(props, { emit, expose }) {
        const affiliation = inject('affiliation');
        const amount = ref('');
        const date = ref('');
        const receipt = ref('');
        const type = ref('');
        const types = ref([]);
        const description = ref('');
        const saveBrotherExpenses = () => {
            fetch('/api/affiliations/expenses', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    affiliation_id: affiliation.value.ID,
                    amount: amount.value,
                    date: date.value,
                    receipt: receipt.value,
                    type: type.value,
                    description: description.value
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
        onMounted(() => {
            fetch('/api/movement-types/manual-expenses-types')
                .then(response => {
                    if (!response.ok) {
                        if (response.status === 401) {
                            window.location.href = '/';
                        }
                    }
                    return response.json();
                })
                .then(data => {
                    types.value = data;
                });
        })
        const toggle = (name) => {
            emit("changeComponent",name);
        }
        return {  amount, date, receipt, description, saveBrotherExpenses, toggle, affiliation, type, types };
        },
    template: `
      <div>
        <h1>Brother Payment</h1>
        <form @submit.prevent="saveBrotherExpenses">
          <div class="mb-3">
            <p>Brother Name</p>
            <p>{{ affiliation.brother.first_name }} {{ affiliation.brother.last_names}}</p>
          </div>
          <div class="mb-3">
            <label for="type" class="form-label">Tipo</label>
            <select id="type" v-model="type" class="form-control">
                <option v-for="type in types" :value="type.name">{{type.description}}</option>
            </select>
            </div>
          
          <div class="mb-3">
            <label for="amount" class="form-label">Cantidad</label>
            <div class="input-group">
              <span class="input-group-text">$</span>
              <input type="number" min="0" step="0.01" data-number-to-fixed="2" data-number-stepfactor="100" class="form-control currency" id="amount"  v-model="amount" />
            </div>      
          </div>
          <div class="mb-3" >
            <label for="date" class="form-label">Date</label>
            <input type="date" id="date" v-model="date" class="form-control date" required />
          </div>
          <div class="mb-3">
            <label for="description" class="form-label">Descripci&oacute;n</label>
            <input type="text" id="description" v-model="description" class="form-control" />
          </div>
          <div class="mb-3">
            <label for="receipt" class="form-label">Recibo</label>
            <input type="text" id="receipt" v-model="receipt" class="form-control" />
          </div>
          <div class="form-footer">
            <div class="mb-3">
            <button type="submit" class="btn btn-primary">Save</button>&nbsp;
            <button type="button" @click="toggle('affiliations')" class="btn btn-secondary">Cancel</button>
          </div>
            </div>
        </form>
      </div>`,
}