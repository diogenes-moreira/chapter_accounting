const { ref, onMounted} = Vue;


export default {
    name: 'ManualMovements',
    setup(props, { emit, expose }) {
        const amount = ref('');
        const date = ref('');
        const receipt = ref('');
        const type = ref('');
        const types = ref([]);
        const description = ref('');
        const saveMovement = () => {
            fetch('/api/chapters/movement', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
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
                });
        }
        onMounted(() => {
            fetch('/api/movement-types/manual-types')
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
        return {  amount, date, receipt, description, saveMovement, toggle, type, types };
    },
    template: `
      <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div class="modal-dialog modal-dialog-scrollable">
          <div class="modal-content">
            <div class="modal-header">
              <p class="modal-title" id="exampleModalLabel"> Crear Movimiento</p>
              <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="saveMovement">
                <div class="mb-3">
                  <label for="type" class="form-label">Tipo</label>
                  <select id="type" v-model="type" class="form-control">
                    <option v-for="type in types" :value="type.name">{{ type.description }}</option>
                  </select>
                </div>
                <div class="mb-3">
                  <label for="amount" class="form-label">Cantidad</label>
                  <div class="input-group">
                    <span class="input-group-text">$</span>
                    <input type="number" min="0" step="0.01" data-number-to-fixed="2" data-number-stepfactor="100"
                           class="form-control currency" id="amount" v-model="amount"/>
                  </div>
                </div>
                <div class="mb-3">
                  <label for="date" class="form-label">Date</label>
                  <input type="date" id="date" v-model="date" class="form-control date" required/>
                </div>
                <div class="mb-3">
                  <label for="description" class="form-label">Descripci&oacute;n</label>
                  <input type="text" id="description" v-model="description" class="form-control"/>
                </div>
                <div class="mb-3">
                  <label for="receipt" class="form-label">Recibo</label>
                  <input type="text" id="receipt" v-model="receipt" class="form-control"/>
                </div>
              </form>
            </div>
            <div class="modal-footer">
            <div class="mb-3">
              <button type="button" @click="saveMovement" class="btn btn-primary">Save</button>&nbsp;
              <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cerrar</button>
            </div>
          </div>
          </div>
        </div>
      </div>`,
}