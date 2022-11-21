<template>
  <div>
    Register
    <v-form>
      <v-text-field
        v-model="fName"
        placeholder="please input your first name"
      />
      <v-text-field
        v-model="lName"
        placeholder="please input your first name"
      />
      <v-text-field
        v-model="tel"
        placeholder="please input your first name"
      />
      <v-text-field
        v-model="email"
        placeholder="please input your first name"
      />

      <v-btn @click="submitForm()" >submit</v-btn>
    </v-form>
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
      this.axios.post("http://localhost:3000/customer/create-profile",data).then((response)=>{
        console.log(response.data)
      })
    }
  }
};
</script>
