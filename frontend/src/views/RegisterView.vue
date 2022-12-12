<template>
  <div class="pa-6 px-lg-16 px-xl-16 mx-lg-auto mx-xl-auto justify-center align-center">
    <v-card class="pa-4 rounded-xl elevation-4">
      <v-card-title class="text-h5 pa-0">Register</v-card-title>
      <v-card-subtitle class="text-wrap pa-0 text-subtitle-1">Create account - a simple form, get the fixing done</v-card-subtitle>
      <v-form class="py-4">
        <v-text-field 
          v-model="fName" 
          box 
          full-width 
          single-line 
          variant="outlined"
          label="First name" 
          background-color="#ffffff" 
          color="blue-darken-3" 
          prepend-inner-icon="mdi-account-outline"
          placeholder="First name"
          />
        <v-text-field 
          v-model="lName" 
          box 
          full-width 
          single-line 
          variant="outlined"
          label="Last name" 
          color="blue-darken-3" 
          prepend-inner-icon="mdi-account-outline"
          placeholder="Last name"
        />
        <v-text-field 
          v-model="tel" 
          box 
          full-width 
          single-line 
          variant="outlined"
          label="Last name" 
          color="blue-darken-3" 
          prepend-inner-icon="mdi-phone-outline"
          placeholder="Last name"
        />
        <v-text-field 
          v-model="email" 
          box 
          full-width 
          single-line 
          variant="outlined"
          label="Last name" 
          color="blue-darken-3" 
          prepend-inner-icon="mdi-email-outline"
          placeholder="Last name"
        />
        <v-btn @click="submitForm()" color="blue-darken-2" block size="large" >Sign up</v-btn>
      </v-form>
    </v-card>
  </div>
</template>

<script>
export default {
  data(){
    return{
      fName: "",
      lName: "",
      tel: "",
      email: ""
    }
  },
  methods:{
    submitForm(){
      const data = new URLSearchParams({
        fName: this.fName,
        lName: this.lName,
        tel: this.tel,
        email: this.email,
      })
      this.axios.post(this.$backendApi + "customer/create-profile",data).then((response)=>{
        console.log(response.data)
        sessionStorage.setItem("auth", true)
        sessionStorage.setItem("cusID", response.data.ID)
        sessionStorage.setItem("fName", response.data.fName)
        sessionStorage.setItem("lName", response.data.lName)
        sessionStorage.setItem("tel", response.data.tel)
        sessionStorage.setItem("email", response.data.email)
        this.$router.push("/")
      })
    }
  }
};
</script>
