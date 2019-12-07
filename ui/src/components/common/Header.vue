<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="info">
      <router-link to="/" class="navbar-brand">Home</router-link>
      <router-link to="/blog" class="navbar-brand">Blog</router-link>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">

          <b-nav-item-dropdown right v-if="info.access_token">
            <!-- Using 'button-content' slot -->
            <template v-slot:button-content>
              <em>User: {{info.user.Email}}</em>
            </template>
            <b-dropdown-item to="/login">Profile</b-dropdown-item>
            <b-dropdown-item to="/login">Sign Out</b-dropdown-item>
          </b-nav-item-dropdown>
          <b-nav-item-dropdown right v-else>
            <template v-slot:button-content>
              <em to="/login">Login</em>
            </template>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
import store from '../../store'
export default {
  name: 'UiHeader',
  data () {
    console.log(store.state)
    return {
      info: store.state
    }
  }
}
</script>

<style scoped>

</style>
