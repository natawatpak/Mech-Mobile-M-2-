<template>
  <div>
    <v-form>
      <InputField
        value="fName"
        placeholder="please input your firstname"
        label=""
        icon=""
        hint=""
        color="#2196F3"
      />
      <InputField
        value="lName"
        placeholder="please input your firstname"
        label=""
        icon=""
        hint=""
        color="#2196F3"
      />

      <v-btn type="submit">submit</v-btn>
    </v-form>
  </div>
</template>

<script>
// @ is an alias to /src
import InputField from "@/components/InputField.vue";

export default {
  components: {
    InputField,
  },
};
</script>
