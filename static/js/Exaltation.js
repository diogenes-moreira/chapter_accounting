const { ref } = Vue;

export default {
    setup(props, { emit }) {
        const firstName = ref('');
        const email = ref('');
        const phoneNumber = ref('');
        const lastNames = ref('');
        const isHonorary = ref(false);
        const toggle = (name) => {
            emit("changeComponent",name);
        }
        const saveExaltation = () => {
            fetch('/api/brothers/exaltation', {
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

        return { firstName, email, phoneNumber, lastNames, isHonorary, saveExaltation, toggle };
    },
    template: `
        <div>
            <h1>Exaltation</h1>
            <form @submit.prevent="saveExaltation">
                <div class="form-group">
                    <label for="firstName">First Name</label>
                    <input type="text" id="firstName" v-model="firstName" class="form-control" required>
                </div>
              <div class="form-group">
                <label for="lastNames">Last Names</label>
                <input type="text" id="lastNames" v-model="lastNames" class="form-control" required>
              </div>
                <div class="form-group">
                    <label for="email">Email</label>
                    <input type="email" id="email" v-model="email" class="form-control" required>
                </div>
                <div class="form-group">
                    <label for="phoneNumber">Phone Number</label>
                    <input type="text" id="phoneNumber" v-model="phoneNumber" class="form-control" required>
                </div>
                <div class="form-group">
                    <label for="isHonorary">Is Honorary</label>
                    <input type="checkbox" id="isHonorary" v-model="isHonorary" class="form-check-input">
                </div>
                <button type="submit" class="btn btn-primary">Save</button>
                <button type="button" class="btn btn-secondary" @click="toggle('affiliations')">Cancel</button>
            </form>
        </div>
    `,
};