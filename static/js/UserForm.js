const { ref, onMounted } = Vue;

export default {
  setup() {
    const user = ref({ user_name: '', password: '', profile: '', chapter_id: '' });
    const users = ref([]);
    const chapters = ref([]);

    const fetchChapters = () => {
        fetch('/api/chapters')
            .then(response => response.json())
            .then(data => {
            chapters.value = data
            });
    };

    const fetchUsers = () => {
      fetch('/api/users')
        .then(response => response.json())
        .then(data => {
          users.value = data;
        });
    };

    const saveUser = () => {
      const method = user.value.ID ? 'PUT' : 'POST';
      const url = user.value.ID ? `/api/users/${user.value.ID}` : '/api/users';

      fetch(url, {
        method: method,
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(user.value)
      }).then(() => {
        fetchUsers();
        user.value = { user_name: '', password: '', profile: '' , chapter_id: ''};
      });
    };

    const editUser = (id) => {
      fetch(`/api/users/${id}`)
        .then(response => response.json())
        .then(data => {
          user.value = data;
        });
    };

    const deleteUser = (id) => {
      fetch(`/api/users/${id}`, {
        method: 'DELETE'
      }).then(() => {
        fetchUsers();
      });
    };

    onMounted(() => {
      fetchUsers();
      fetchChapters();
    });

    return {
      user,
      users,
      chapters,
      saveUser,
      editUser,
      deleteUser,
    };
  },
  methods: {
    chapterName(id) {
      try {
        return this.chapters.find(chapter => chapter.ID === id).name;
      }catch (e){
        return "";
      }
    }
  },
    template:`<div>
      <h2>User CRUD</h2>
      <form @submit.prevent="saveUser" class="form">
        <input type="hidden" v-model="user.id">
        <div class="mb-3">
          <label for="user_name" class="form-label">User Name</label>
          <input type="text" v-model="user.user_name" required class="form-control">
        </div>
        <div class="mb-3">
          <label for="password" class="form-label">Password</label>
          <input type="password" v-model="user.password" class="form-control">
        </div>
        <div class="mb-3">
          <label for="profile" class="form-label">Profile</label>
          <select v-model="user.profile" class="form-control form-select">
            <option value="admin">Administrador</option>
            <option value="treasurer">Tesorero</option>
            <option value="principal">Principal</option>
            <option value="companion">Compa&ntilde;ero</option>
          </select>
        </div>
        <div class="mb-3">
          <label for="profile" class="form-label">Capitulo</label>
          <select v-model="user.chapter_id" class="form-control form-select">
            <option v-for="chapter in chapters" :key="chapter.ID" :value="chapter.ID">{{ chapter.name }}</option>
          </select>
        </div>
        <button type="submit" class="btn btn-primary">Save</button>
      </form>
      <h2>User List</h2>
      <div class="m-3">
      <table class="table table-bordered">
        <thead>
            <tr>
                <th>User Name</th>
                <th>Profile</th>
                <th>Chapter</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="user in users" :key="user.ID">
                <td>{{ user.user_name }}</td>
                <td>{{ user.profile }}</td>
                <td>{{ chapterName(user.chapter_id) }}</td>
                <td>
                    <button @click="editUser(user.ID)" class="btn btn-warning">Edit</button>
                    <button @click="deleteUser(user.ID)" class="btn btn-danger">Delete</button>
                </td>
            </tr>
        </tbody>
        </table>
      </div>
    </div>`
};