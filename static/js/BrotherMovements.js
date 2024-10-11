const {  inject } = Vue;

export default {
    name: 'BrotherMovements',
    setup(props) {
        const affiliation = inject('affiliation');
        return { affiliation };
    },
    methods: {
        dateFormated(date) {
            return new Date(date).toLocaleDateString();
        }
    },
    template: `
      <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div class="modal-dialog modal-dialog-scrollable">
          <div class="modal-content">
            <div class="modal-header">
              <p class="modal-title" id="exampleModalLabel">Movimientos</p>
              <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
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
                  <tr v-for="movement in affiliation.movements">
                    <td>{{ dateFormated(movement.date) }}</td>
                    <td>{{ movement.description }}</td>
                    <td>{{ movement.movement_type.description }}</td>
                    <td v-if="movement.movement_type.credit">{{ movement.amount }}</td>
                    <td v-else></td>
                    <td v-if="!movement.movement_type.credit">{{ movement.amount }}</td>
                    <td v-else></td>
                    <td>{{ movement.receipt }}</td>
                  </tr>
                  </tbody></table>
              </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cerrar</button>
            </div>
          </div>
        </div>
        </div>
    `


}