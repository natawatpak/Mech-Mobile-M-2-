<template>
  <div class="text-center mx-16 justify-center">
    <v-row class="pa-4 ma-4 align-center">
      <router-link to="/order" class="text-decoration-none">
      <v-btn
        icon
        class="hidden-xs-only"
      >
        <v-icon>mdi-arrow-left</v-icon>
      </v-btn>
    </router-link>
      <v-toolbar-title class="text-left text-h4 px-6">Order {{details.id}}</v-toolbar-title>
      <v-chip color="yellow-darken-3" text-color="white">{{details.status}}</v-chip>
    </v-row>

    <v-card class="text-left mx-8 mb-4 pa-4" variant="tonal">
      <v-card-title class="text-h6">Status</v-card-title>
      <v-timeline direction="horizontal" line-inset="12" class="pa-4">
      <v-timeline-item size="large" v-model="items" :dot-color="currentState == 1 || currentState == 2 || currentState == 3 || currentState == 4? 'green':'white'" :icon=items[0]>
        <template v-slot:opposite>
          Accept
        </template>
       </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="currentState == 2 || currentState == 3 || currentState == 4? 'green':'white'" :icon=items[1]>
      <template v-slot:opposite>
      </template>
        On the way
      </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="currentState == 3 || currentState == 4? 'green':'white'" :icon=items[2]>
        <template v-slot:opposite>
          On process
        </template>
      </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="currentState == 4? 'green':'white'" :icon=items[3]>
        <template v-slot:opposite>
        </template>
        Finish
      </v-timeline-item>
      </v-timeline>
      <v-card-subtitle v-if="currentState == 1">{{otw}}</v-card-subtitle>
      <v-card-subtitle v-else-if="currentState == 2">{{onp}}</v-card-subtitle>
      <v-card-subtitle v-else-if="currentState == 3">{{fin}}</v-card-subtitle>
      <v-row justify="end" class="pa-4">
        <v-btn :disabled="valid" class="mx-1 " variant="tonal" color="blue" @click="checkStage">Next stage</v-btn>
      </v-row>
    </v-card>

    <v-card class="text-left mx-8 mb-4 pa-4" variant="tonal">
      <v-title class="text-h6">Car Details</v-title>
        <v-card-text class="text-h7">
          Type: {{details.type}} | Brand: {{details.brand}}
          <br />
          License plate: {{details.plate}}
        </v-card-text>
    </v-card>

    <v-card variant="tonal" class="text-left mx-8 mb-4 pa-4 d-flex justify-left align-center">
      <section>
        <v-card-title class="text-h6">Current Location</v-card-title>
        <v-card-text>{{details.lat+', '+ details.lng}}</v-card-text>
      </section>
    </v-card>

    <v-card class="text-left mx-8 mb-4 pa-4" variant="tonal">
      <v-title class="text-h6">Problems</v-title>
      <v-card-text class="text-h7">
        <p v-for="p in problems" :key="p">{{'-' + p}}</p>
      </v-card-text>
    </v-card>

    <v-card class="text-left mx-8 mb-4 pa-4" variant="tonal">
      <v-title class="text-h6">Description</v-title>
      <v-card-text class="text-h7">
        <p>{{description}}</p>
      </v-card-text>
    </v-card> 

    <v-btn class="mx-1" color="error" variant="tonal" @click="dialog3 = true">Cancel</v-btn>

    <v-dialog v-model="dialog3" width="1000" class="text-left pa-4">
      <v-card class="text-left pa-4">
        <v-card-title class="text-h5">Cancel order</v-card-title>  
        <v-card-text>
          <p>Why are you cancelling?</p>   
          <v-radio-group>
            <v-radio label="The problem was already fixed" value="fixed"></v-radio>
            <v-radio label="Client want to cancel" value="cancel"></v-radio>
            <v-radio label="No need to fix anymore" value="no need"></v-radio>
            <v-radio label="Others" value="other"></v-radio>
          </v-radio-group>
        </v-card-text>
        <v-row justify="end" class="pa-4">
          <v-card-action>
            <v-btn class="mx-1" variant="tonal" color="error" @click="dialog3 = false">
              Cancel
            </v-btn>

            <v-btn class="mx-1" color="blue" variant="tonal" @click="submitForm()">
              Submit
            </v-btn>
          </v-card-action>
        </v-row>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialog" width="800" class="text-left pa-4">
          <v-card class="text-left pa-4">
          <v-card-title class="text-h5">Status {{currentState}}</v-card-title> 
          <v-timeline direction="horizontal" line-inset="12" class="pa-4">
            <v-timeline-item size="large" v-model="items" :dot-color="currentState == '1' || currentState == '2' || currentState == '3' || currentState == '4'? 'green':'white'" :icon=items[0]>
            <template v-slot:opposite>
              Accept
            </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="currentState == '2' || currentState == '3' || currentState == '4'? 'green':'white'" :icon=items[1]>
            <template v-slot:opposite>
            </template>
              On the way
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="currentState == '3' || currentState == '4'? 'green':'white'" :icon=items[2]>
            <template v-slot:opposite>
              On process
            </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="currentState == '4'? 'green':'white'" :icon=items[3]>
            <template v-slot:opposite>
            </template>
              Finish
            </v-timeline-item>
            </v-timeline>
            
            <div class="px-2">
              <v-text class="text-h6 pa-0">Please type 'confirm' to continue</v-text>
              <v-text-field v-model="confirm" label="Confirmation" required></v-text-field>
            </div>

            <v-row justify="end" class="pa-4">
                <v-card-action>
                  <v-btn class="mx-1" variant="tonal" color="error" @click="dialog = false">
                    Cancel
                  </v-btn>

                  <v-btn :disabled="valid2" class="mx-1" color="blue" variant="tonal" @click="checkConfirm()">
                    Confirm
                  </v-btn>
                </v-card-action>
              </v-row>
          </v-card>
        </v-dialog>

        <v-dialog v-model="dialog2" width="800" class="text-left pa-4">
          <v-card class="text-left pa-4">
          <v-card-title class="text-h5">Status {{currentState}}</v-card-title> 
          <v-timeline direction="horizontal" line-inset="12" class="pa-4">
            <v-timeline-item size="large" v-model="items" :dot-color="currentState == '1' || currentState == '2' || currentState == '3' || currentState == '4'? 'green':'white'" :icon=items[0]>
            <template v-slot:opposite>
              Accept
            </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="currentState == '2' || currentState == '3' || currentState == '4'? 'green':'white'" :icon=items[1]>
            <template v-slot:opposite>
            </template>
              On the way
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="currentState == '3' || currentState == '4'? 'green':'white'" :icon=items[2]>
            <template v-slot:opposite>
              On process
            </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="currentState == '4'? 'green':'white'" :icon=items[3]>
            <template v-slot:opposite>
            </template>
              Finish
            </v-timeline-item>
            </v-timeline>

            <v-card-text>
              <v-radio-group column>
                <v-radio label="All complete" value="complete"></v-radio>
                <v-radio label="Go to garage" value="garage"></v-radio>
              </v-radio-group>
            </v-card-text>

            <div class="px-2">
              <v-text class="text-h6 pa-0">Please type 'confirm' to continue</v-text>
              <v-text-field v-model="confirm" label="Confirmation" required></v-text-field>
            </div>
            <v-row justify="end" class="pa-4">
                <v-card-action>
                  <v-btn class="mx-1" variant="tonal" color="error" @click="dialog2 = false">
                    Cancel
                  </v-btn>

                  <v-btn class="mx-1" color="blue" variant="tonal" @click="checkConfirm()">
                    Submit
                  </v-btn>
                </v-card-action>
              </v-row>
          </v-card>
        </v-dialog>
        
  </div>
</template>

<style>

</style>

<script>
export default {
  data() {
    return {
      ticketID: undefined,
      valid: false,
      valid2: false,
      dialog: false,
      dialog2: false,
      dialog3: false,
      currentState: 2,
      confirm: '',
      description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam rutrum tincidunt arcu sed gravida. Suspendisse mollis ex sed magna viverra, eu tincidunt velit lobortis. Phasellus accumsan mauris est, in pretium orci lacinia in. Nulla in bibendum ex, eu condimentum mauris. Pellentesque quis nisl a justo ultrices molestie ac mattis libero. Aliquam vel sollicitudin quam, vel molestie nunc. Proin dolor dolor, vehicula sed nisl dapibus, semper scelerisque nisl. Sed id neque ac metus efficitur vestibulum. Phasellus quis nibh ac dui molestie fringilla sit amet bibendum sem. Quisque consectetur sit amet nunc ac elementum. Suspendisse vulputate mauris erat, nec tempus dui vulputate.",
      problems: ["no battery", "broken motor"],
      details: {
        id: "12345",
        type: "SUV",
        brand: "MG",
        plate: "ศง 4727",
        name: "nunnapat",
        problem: "2",
        location: "2.1",
        lat: "20",
        lng: "90",
        status: "Accept",
      },
      socket: undefined,
      items: ['mdi-thumb-up', 'mdi-car', 'mdi-wrench', 'mdi-check'],
      otw: "Next stage: 'On the way'. Are your mechanic ready to go?",
      onp: "Next stage: 'On process'. Your mechanic is fixing customer car.",
      fin: "Last stage: 'Finish'. All completed!",
    };
  },
  methods: {
    checkConfirm() {
      if (this.confirm == 'confirm') {
        if (this.currentState == 1) {
          this.dialog = false;
          this.confirm = '';
          this.currentState++;
          this.details.status = 'On the way';
          this.updateTicketStatus('On the way')
        }
        else if (this.currentState == 2) {
          this.dialog = false;
          this.confirm = '';
          this.currentState++;
          this.details.status = 'In process';
          this.updateTicketStatus('In process')

        }
        else if (this.currentState == 3) {
          this.dialog2 = false;
          this.confirm = '';
          this.currentState++;
          this.details.status = 'Complete';
          this.updateTicketStatus('Complete')

        }
      }
      // else {
      //   this.valid2 = true;
      // }
    },
    checkStage() {
      if (this.currentState == 1) {
        this.dialog = true;      
      }
      else if (this.currentState == 2) {
        this.dialog = true;
      }
      else if (this.currentState == 3) {
        this.dialog2 = true;
        // this.valid = true;
      }
      else if (this.currentState == 4) {
        this.valid = true;
      }
    },
    updateTicketStatus(status){
      this.socket.send(status);

      const data = new URLSearchParams({
        ticketID: this.ticketID,
        cusID: this.cusID,
        carID: this.carID,
        problem: this.problem,
        shopID: sessionStorage.getItem("shopID"),
        status: status,
        // lat:
        // lng:
        // description:
      });
      this.axios
        .post("http://localhost:3000/shop/update-ticket", data)
        .then((response) => {
          console.log(response.data);
          this.cusID = response.data.cusID;
          this.carID = response.data.carID;
          this.shopID = response.data.shopID;
          this.status = response.data.status;
          this.problems = response.data.problem;
        })
    },
    getTicket(){
      const data = new URLSearchParams({
        ticketID: this.ticketID,
      });
      this.axios
        .post("http://localhost:3000/customer/get-ticket", data)
        .then((response) => {
          console.log(response.data);
          this.carID = response.data.carID;
          this.shopID = response.data.shopID;
          this.status = response.data.status;
          this.problems = response.data.problem;
          this.createdTime = response.data.createdTime;
          this.acceptedTime = response.data.acceptedTime;
          this.getCarDetail()
        })
    },
    getCarDetail(){
      const data = new URLSearchParams({
        carID: this.carID,
      });
      this.axios
        .post("http://localhost:3000/customer/get-car", data)
        .then((response) => {
          this.car = response.data;
        });
    },
  },
  mounted() {
    this.ticketID = sessionStorage.getItem("ticketID")
    console.log(sessionStorage.getItem("ticketID"))

    this.socket = new WebSocket("ws://127.0.0.1:3000/shop/ws/"+sessionStorage.getItem("ticketID"));
    console.log("Attempting Connection...");

    this.socket.onopen = () => {
      console.log("Successfully Connected");
      this.socket.send("Accepted");
    };

    this.socket.addEventListener("message", (event) => {
      console.log("Message from server ", event.data);
    });

    this.socket.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
      this.socket.send("Client Closed!");
    };

    this.socket.onerror = (error) => {
      this.console.log("Socket Error: ", error);
    };

  },
    submitForm(){
      this.dialog3 = false
      const data = new URLSearchParams({
        reason: this.value,
      })
      this.axios.post("http://localhost:3000/customer/create-profile",data).then((response)=>{
        console.log(response.data)
      })
    },
};
</script>