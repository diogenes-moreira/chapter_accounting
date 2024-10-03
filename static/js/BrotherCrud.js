Vue.component('chapter-crud', {
    data() {
        return {
            chapters: [],
            newChapter: { name: '' },
            editChapter: null
        };
    },
    mounted() {
        this.fetchChapters();
    },
    methods: {
        fetchChapters() {
            fetch('/brothers')
                .then(response => response.json())
                .then(data => {
                    this.brothers = data;
                });
        },
        createBrother() {
            fetch('/brothers', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(this.newBrother)
            })
                .then(response => response.json())
                .then(data => {
                    this.brothers.push(data);
                    this.newBrother = { name: ''};
                });
        },
        updateBrother(brother) {
            fetch(`/brothers/${brother.ID}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(brother)
            })
                .then(response => response.json())
                .then(data => {
                    this.editBrother = null;
                    this.fetchBrothers();
                });
        },
        deleteBrother(id) {
            fetch(`/brothers/${id}`, { method: 'DELETE' })
                .then(() => {
                    this.brothers = this.brothers.filter(brother => brother.id !== id);
                });
            this.fetchBrothers()
        }
    },
    template: `
        <div>
            <h2>Create Brother</h2>
            <form @submit.prevent="createBrother">
                <div class="form-group">
                    <label for="Name">Name</label>
                    <input type="text" v-model="newBrother.name" class="form-control" id="title" required>
                </div>
                <button type="submit" class="btn btn-primary">Create</button>
            </form>

            <h2 class="mt-5">Brothers</h2>
            <ul class="list-group">
                <li v-for="brother in brothers" :key="brother.id" class="list-group-item">
                    <div v-if="editBrother && editBrother.id === brother.id">
                        <input v-model="editBrother.name" class="form-control mb-2">
                        <button @click="updateBrother(editBrother)" class="btn btn-success">Save</button>
                        <button @click="editBrother = null" class="btn btn-secondary">Cancel</button>
                    </div>
                    <div v-else>
                        <h5>{{ brother.name }}</h5>
                        <button @click="editBrother = { ...brother }" class="btn btn-warning">Edit</button>
                        <button @click="deleteBrother(brother.ID)" class="btn btn-danger">Delete</button>
                    </div>
                </li>
            </ul>
        </div>
    `
});

new Vue({
    el: '#app'
});