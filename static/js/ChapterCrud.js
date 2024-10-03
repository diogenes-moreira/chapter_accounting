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
            fetch('/chapters')
                .then(response => response.json())
                .then(data => {
                    this.chapters = data;
                });
        },
        createChapter() {
            fetch('/chapters', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(this.newChapter)
            })
                .then(response => response.json())
                .then(data => {
                    this.chapters.push(data);
                    this.newChapter = { name: ''};
                });
        },
        updateChapter(chapter) {
            fetch(`/chapters/${chapter.ID}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(chapter)
            })
                .then(response => response.json())
                .then(data => {
                    this.editChapter = null;
                    this.fetchChapters();
                });
        },
        deleteChapter(id) {
            fetch(`/chapters/${id}`, { method: 'DELETE' })
                .then(() => {
                    this.chapters = this.chapters.filter(chapter => chapter.id !== id);
                });
            this.fetchChapters()
        }
    },
    template: `
        <div>
            <h2>Create Chapter</h2>
            <form @submit.prevent="createChapter">
                <div class="form-group">
                    <label for="Name">Name</label>
                    <input type="text" v-model="newChapter.name" class="form-control" id="title" required>
                </div>
                <button type="submit" class="btn btn-primary">Create</button>
            </form>

            <h2 class="mt-5">Chapters</h2>
            <ul class="list-group">
                <li v-for="chapter in chapters" :key="chapter.id" class="list-group-item">
                    <div v-if="editChapter && editChapter.id === chapter.id">
                        <input v-model="editChapter.name" class="form-control mb-2">
                        <button @click="updateChapter(editChapter)" class="btn btn-success">Save</button>
                        <button @click="editChapter = null" class="btn btn-secondary">Cancel</button>
                    </div>
                    <div v-else>
                        <h5>{{ chapter.name }}</h5>
                        <button @click="editChapter = { ...chapter }" class="btn btn-warning">Edit</button>
                        <button @click="deleteChapter(chapter.ID)" class="btn btn-danger">Delete</button>
                    </div>
                </li>
            </ul>
        </div>
    `
});

new Vue({
    el: '#app'
});