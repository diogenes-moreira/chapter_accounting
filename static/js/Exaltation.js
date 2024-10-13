const { ref, inject } = Vue;

export default {
    setup(props, { emit }) {
        const  exaltation = inject('exaltation');
        const firstName = ref('');
        const email = ref('');
        const phoneNumber = ref('');
        const lastNames = ref('');
        const isHonorary = ref(false);
        const toggle = (name) => {
            emit("changeComponent",name);
        }
        const saveExaltation = () => {
            const url = exaltation.value ? '/api/brothers/exaltation' : '/api/brothers/affiliation';
            fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    brother: {
                        first_name: firstName.value,
                        email: email.value,
                        phone_number: phoneNumber.value,
                        last_names: lastNames.value
                    },
                    is_honorary: isHonorary.value
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
                    toggle('affiliations');
                    console.log(data);
                });
        };

        return { firstName, email, phoneNumber, lastNames, isHonorary, exaltation, saveExaltation, toggle };
    },
    template: `
        <div >
            <h1 v-if="exaltation">Exaltation</h1>
            <h1 v-else>Affiliation</h1>
            <form @submit.prevent="saveExaltation">
                <div class="mb-3">
                    <label class="form-label" for="firstName">First Name</label>
                    <input type="text" id="firstName" v-model="firstName" class="form-control" required>
                </div>
              <div class="mb-3">
                <label class="form-label" for="lastNames">Last Names</label>
                <input type="text" id="lastNames" v-model="lastNames" class="form-control" required>
              </div>
                <div class="mb-3">
                    <label class="form-label" for="email">Email</label>
                    <input type="email" id="email" v-model="email" class="form-control" required>
                </div>
                <div class="mb-3">
                    <label class="form-label" for="phoneNumber">Phone Number</label>
                    <input type="text" id="phoneNumber" v-model="phoneNumber" class="form-control" required>
                </div>
                <div class="mb-3">
                  <div class="form-check">
                    <input type="checkbox" id="isHonorary" v-model="isHonorary" class="form-check-input">
                    <label for="isHonorary" class="form-check-label">Is Honorary</label>
                  </div>
                </div>
                
              <div class="mb-3">
                <div class="form-group">
                <button type="submit" class="btn btn-primary">Save</button>&nbsp;
                <button type="button" @click="toggle('affiliations')" class="btn btn-secondary">Cancel</button>
              </div>
            </div>
            </form>
        </div>
    `,
};