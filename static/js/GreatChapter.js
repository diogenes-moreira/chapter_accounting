const {ref, provide, onMounted} = Vue;

export default {
    setup: function () {
        const pendingInstallments = ref([]);
        const dueInstallments = ref([]);
        const deposits = ref([]);
        const balance = ref(0);
        const pendingAmount = ref(0);
        const dueAmount = ref(0);
        const totalDeposits = ref(0);
        const installments = ref([]);
        let Format = new Intl.NumberFormat('es-AR',
            {style: 'currency', currency: 'ARS', minimumFractionDigits: 2, currencySign: 'accounting'});
        const showDetail = (movement) => {
            installments.value = movement.installments;
            for (let installment of installments.value) {
                installment.date = new Date(installment.paid_date).toLocaleDateString();
                installment.great_chapter_amount = Format.format(installment.great_chapter_amount);
            }
        }
        const fetchGreatChapter = () => {
            fetch('/api/chapters/great-chapter')
                .then(response => {
                    if (!response.ok) {
                        if (response.status === 401) {
                            window.location.href = '/';
                        }
                    }
                    return response.json();
                })
                .then(data => {
                    pendingInstallments.value = []
                    dueInstallments.value = []
                    deposits.value = []
                    balance.value = Format.format(data.balance);
                    pendingAmount.value = Format.format(data.pending_amount);
                    dueAmount.value = Format.format(data.due_amount);
                    totalDeposits.value = Format.format(data.total_deposits);
                    for (let installment of data.pending_installments) {
                        installment.date = new Date(installment.paid_date).toLocaleDateString();
                        installment.great_chapter_amount = Format.format(installment.great_chapter_amount);
                        pendingInstallments.value.push(installment);
                    }
                    for (let installment of data.due_installments) {
                        installment.date = new Date(installment.due_date).toLocaleDateString();
                        installment.great_chapter_amount = Format.format(installment.great_chapter_amount);
                        dueInstallments.value.push(installment);
                    }
                    for (let deposit of data.deposits) {
                        deposit.date = new Date(deposit.deposit_date).toLocaleDateString();
                        deposit.amount = Format.format(deposit.amount);
                        deposits.value.push(deposit);
                    }
                });
        };
        onMounted(() => {
            fetchGreatChapter();
        });
        provide('fetchGreatChapter', fetchGreatChapter);
        const generateDeposit = () => {
            let selected = pendingInstallments.value.filter(installment => installment.selected);
            let ids = selected.map(installment => installment.ID);
            if (ids.length === 0) {
                alert('Debe seleccionar al menos una cuota');
                return;
            }
           if(!confirm('¿Desea generar un deposito con las cuotas seleccionadas?')){
               return;
           }
            fetch('/api/chapters/deposit', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    installments: ids
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
                    fetchGreatChapter();
        })};

        return {
            pendingInstallments,
            balance,
            pendingAmount,
            dueAmount,
            dueInstallments,
            deposits,
            totalDeposits,
            installments,
            fetchGreatChapter,
            generateDeposit,
            showDetail
        };
    },
    template: `
      <div>
        <h2>Gran Capitulo</h2>
        <div class="row">
          <div class="col">
            <b>Pendientes de Deposito</b>
            <p>{{ pendingAmount }}</p>
          </div>
          <div class="col">
            <b>Total Pendiente Gran Capitulo</b>
            <p>{{ balance }}</p>
          </div>
        </div>
        <a role="button" @click="generateDeposit"><i class="bi bi-plus"></i>Generar Deposito</a>&nbsp;&nbsp;
        <table class="table">
          <thead>
          <tr>
            <th></th>
            <th>Fecha de Pago</th>
            <th>Descripci&oacute;n</th>
            <th>Monto</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="movement in pendingInstallments">
            <td><input type="checkbox" v-model="movement.selected" /></td>
            <td>{{ movement.date }}</td>
            <td>{{ movement.description }} - {{ movement.brother }}</td>
            <td>{{ movement.great_chapter_amount }}</td>
          </tr>
          </tbody>
        </table>
        <div class="row">
          <div class="col">
            <b>Cuotas Vencidas</b>
            <p>{{ dueAmount }}</p>
          </div>
        </div>
        <table class="table">
          <thead>
          <tr>
            <th>Fecha de Vencimiento</th>
            <th>Descripci&oacute;n</th>
            <th>Monto</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="movement in dueInstallments">
            <td>{{ movement.date }}</td>
            <td>{{ movement.description }} - {{ movement.brother }}</td>
            <td>{{ movement.great_chapter_amount }}</td>
          </tr>
          </tbody>
        </table>
        <div class="row">
          <div class="col">
            <b>Depositos</b>
            <p>{{ totalDeposits }}</p>
          </div>
        </div>
        <table class="table">
          <thead>
          <tr>
            <th>Fecha de Deposito</th>
            <th>Monto</th>
            <th></th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="movement in deposits">
            <td>{{ movement.date }}</td>
            <td>{{ movement.amount }}</td>
            <td><a role="button" @click="showDetail(movement)" data-bs-target="#detailModal" data-bs-toggle="modal"><i class="bi bi-zoom-in"></i></a></td>
          </tr>
          </tbody>
        </table>
        <div class="modal fade" id="detailModal" tabindex="-1" aria-labelledby="detailModalLabel" aria-hidden="true">
            <div class="modal-dialog modal-xl">
                <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="detailModalLabel">Detalle de Cuotas</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <table class="table">
                    <thead>
                    <tr>
                        <th>Fecha de Pago</th>
                        <th>Descripci&oacute;n</th>
                        <th>Monto</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr v-for="installment in installments">
                        <td>{{ installment.date }}</td>
                        <td>{{ installment.description }} - {{ installment.brother }}</td>
                        <td>{{ installment.great_chapter_amount }}</td>
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
      </div>`
};
