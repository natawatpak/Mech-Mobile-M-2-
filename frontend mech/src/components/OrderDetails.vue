<template>
  <div class="text-center ma-4 justify-center">
    <v-row class="ma-4 justify-space-between align-center" v-for="item in incoming" :key="item.id">
      <div class="text-left text-h4 my-2">{{item.cus.fName}} | {{item.car.plate}}</div>
      <v-chip color="yellow-darken-3" text-color="white" v-if="(item.status == stage[0]) || (item.status == stage[1]) || (item.status == stage[2])">{{item.status}}</v-chip>
      <v-chip color="green" text-color="white" v-if="(item.status == stage[3] || incoming.status == stage[4])">{{item.status}}</v-chip>
    </v-row>

    <v-card class="text-left mb-4 pa-4 rounded-lg elevation-4">
      <v-card-title class="text-h6 pa-0">Status</v-card-title>
      <v-timeline direction="horizontal" line-inset="12" class="pa-0" v-for="item in incoming" :key="item.id">
      <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[0] || item.status == stage[1] || item.status == stage[2] || item.status == stage[3] || item.status == stage[4] ? 'green-lighten-1':'white'" :icon=items[0]>
        <template v-slot:opposite>
          Accept
        </template>
       </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[1] || item.status == stage[2] || item.status == stage[3] || item.status == stage[4]? 'green-lighten-1':'white'" :icon=items[1]>
      <template v-slot:opposite>
      </template>
        On the way
      </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[2] || item.status == stage[3] || item.status == stage[4]? 'green-lighten-1':'white'" :icon=items[2]>
        <template v-slot:opposite>
          On process
        </template>
      </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[3] || item.status == stage[4]? 'green-lighten-1':'white'" :icon=items[3]>
        <template v-slot:opposite>
        </template>
        Finish
      </v-timeline-item>
      </v-timeline>
      <div  v-for="item in incoming" :key="item.id">
      <v-card-subtitle v-if="(item.status == stage[0])" class="pa-0 py-2 text-wrap">{{otw}}</v-card-subtitle>
      <v-card-subtitle v-else-if="(item.status == stage[1])" class="pa-0 py-2 text-wrap">{{onp}}</v-card-subtitle>
      <v-card-subtitle v-else-if="(item.status == stage[2])" class="pa-0 py-2 text-wrap">{{fin}}</v-card-subtitle>
      </div>
      <v-row justify="end" class="pa-4">
        <v-btn :disabled="valid" class="mx-1 " variant="tonal" color="blue-darken-2" @click="checkStage">Next stage</v-btn>
      </v-row>
    </v-card>

    <v-card v-for="item in incoming" :key="item.id" class="text-left mb-4 pa-4 d-flex justify-left align-center rounded-lg elevation-4">
      <section>
        <v-card-title class="text-h6 pa-0">Customer name</v-card-title>
        <v-card-text class="pa-0 py-2" v-for="item in incoming" :key="item.id">
          {{(ticket.cus.fName+' '+ticket.cus.lName)}}
        </v-card-text>
      </section>
    </v-card>

    <v-card class="text-left mb-4 pa-4 rounded-lg elevation-4">
      <v-title class="text-h6 pa-0">Car Details</v-title>
        <v-card-text class="text-h7 pa-0 py-2" v-for="item in incoming" :key="item.id">
          Type: {{item.car.type}}
          <br>
          Brand: {{item.car.brand}}
          <br />
          License plate: {{item.car.plate}}
        </v-card-text>
    </v-card>

    

    <v-card v-for="item in incoming" :key="item.id" class="text-left mb-4 pa-4 d-flex justify-left align-center rounded-lg elevation-4">
      <section>
        <v-card-title class="text-h6 pa-0">Current Location</v-card-title>
        <v-card-text class="pa-0 py-2" v-for="item in incoming" :key="item.id">
          {{item.location.lat+', '+ item.location.lng}} 
          <br>
          Distance: {{item.location.distance}} km
        </v-card-text>
      </section>
    </v-card>

    <v-card class="text-left mb-4 pa-4 rounded-lg elevation-4">
      <v-title class="text-h6 pa-0">Problems</v-title>
      <v-card-text class="text-h7 pa-0">
        <p v-for="p in ticket.problem" :key="p">{{'- ' + p}}</p>
      </v-card-text>
    </v-card>

    <v-card class="text-left mb-4 pa-4 rounded-lg elevation-4">
      <v-title class="text-h6 pa-0">Description</v-title>
      <v-card-text class="text-h7 pa-0">
        <p>{{ticket.description}}</p>
      </v-card-text>
    </v-card> 

    <div class="text-center pa-4">
        <v-card-action>
          <v-btn :disabled="valid" color="red" variant="tonal" @click="checkStage()" v-if="(ticket.status==stage[0] || ticket.status==stage[1]|| ticket.status==stage[2])">Cancel</v-btn>
          <v-btn color="blue-darken-2" variant="tonal" to="/" v-if="(ticket.status== stage[3] || ticket.status==stage[4])">Go home</v-btn>
        </v-card-action>
    </div>

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
          <v-card-title class="text-h5">Status {{incoming.status}}</v-card-title> 
          <v-timeline direction="horizontal" line-inset="12" class="pa-0" v-for="item in incoming" :key="item.id">
            <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[0] || item.status == stage[1] || item.status == stage[2] || item.status == stage[3] || item.status == stage[4] ? 'green-lighten-1':'white'" :icon=items[0]>
              <template v-slot:opposite>
                Accept
              </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[1] || item.status == stage[2] || item.status == stage[3] || item.status == stage[4]? 'green-lighten-1':'white'" :icon=items[1]>
            <template v-slot:opposite>
            </template>
              On the way
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[2] || item.status == stage[3] || item.status == stage[4]? 'green-lighten-1':'white'" :icon=items[2]>
              <template v-slot:opposite>
                On process
              </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[3] || item.status == stage[4]? 'green-lighten-1':'white'" :icon=items[3]>
              <template v-slot:opposite>
              </template>
              Finish
            </v-timeline-item>
            </v-timeline>
            
            <div class="pa-2">
              <v-text class="text-subtitle-1 pa-0">Please type 'confirm' to continue</v-text>
              <v-text-field v-model="confirm" label="Confirm" required></v-text-field>
            </div>

            <v-row justify="end" class="pa-4">
                <v-card-action>
                  <v-btn class="mx-1" variant="tonal" color="error" @click="dialog = false">
                    Cancel
                  </v-btn>

                  <v-btn :disabled="valid2" class="mx-1" color="blue-darken-2" variant="tonal" @click="checkConfirm()">
                    Confirm
                  </v-btn>
                </v-card-action>
              </v-row>
          </v-card>
        </v-dialog>

        <v-dialog v-model="dialog2" width="800" class="text-left pa-4">
          <v-card class="text-left pa-4">
          <v-card-title class="text-h5">Status {{incoming.status}}</v-card-title> 
          <v-timeline direction="horizontal" line-inset="12" class="pa-0" v-for="item in incoming" :key="item.id">
            <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[0] || item.status == stage[1] || item.status == stage[2] || item.status == stage[3] || item.status == stage[4] ? 'green-lighten-1':'white'" :icon=items[0]>
              <template v-slot:opposite>
                Accept
              </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[1] || item.status == stage[2] || item.status == stage[3] || item.status == stage[4]? 'green-lighten-1':'white'" :icon=items[1]>
            <template v-slot:opposite>
            </template>
              On the way
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[2] || item.status == stage[3] || item.status == stage[4]? 'green-lighten-1':'white'" :icon=items[2]>
              <template v-slot:opposite>
                On process
              </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="item.status == stage[3] || item.status == stage[4]? 'green-lighten-1':'white'" :icon=items[3]>
              <template v-slot:opposite>
              </template>
              Finish
            </v-timeline-item>
            </v-timeline>

            <v-card-text>
              <v-radio-group column v-model="finishOption">
                <v-radio label="All complete" value="Finish:Completed"></v-radio>
                <v-radio label="Go to garage" value="Finish:Garage"></v-radio>
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
      finishOption: null,
      stage: [ 'Accepted', 'On the way', 'Processing', 'Finish:Garage', 'Finish:Completed'],
      confirm: '',
      ticket: "status: 'Accepted'",
      items: ['mdi-thumb-up', 'mdi-car', 'mdi-wrench', 'mdi-check'],
      otw: "Next stage: 'On the way'. Are your mechanic ready to go?",
      onp: "Next stage: 'On process'. Your mechanic is fixing customer car.",
      fin: "Last stage: 'Finish'. All completed!",
    };
  },
  methods: {
    checkConfirm() {
      if (this.confirm == 'confirm') {
        if (this.ticket.status == this.stage[0]) {
          this.dialog = false;
          this.confirm = '';
          this.currentState++;
          // this.details.status = 'On the way';
          this.updateTicketStatus('On the way')
          this.ticket.status = 'On the way';
        }
        else if (this.ticket.status == this.stage[1]) {
          this.dialog = false;
          this.confirm = '';
          this.currentState++;
          // this.details.status = 'In process';
          this.updateTicketStatus('In process')
          this.ticket.status = 'Processing';
        }
        else if (this.ticket.status == this.stage[2]) {
          this.dialog2 = false;
          this.confirm = '';
          this.currentState++;
          // this.details.status = 'Complete';
          this.updateTicketStatus('Complete')
          this.ticket.status = this.finishOption;
        }
      }
    },
    checkStage() {
      if (this.ticket.status == this.stage[0] || this.ticket.status == this.stage[1]) {
        this.dialog = true;      
      }
      else if (this.ticket.status == this.stage[2]) {
        this.dialog2 = true;
      }
      else {
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
        .post(this.$backendApi+"shop/update-ticket", data)
        .then((response) => {
          console.log(response.data);
          this.cusID = response.data.cusID;
          this.carID = response.data.carID;
          this.shopID = response.data.shopID;
          this.status = response.data.status;
          this.problems = response.data.problem;
        })
    },
    getTicket(id) {
      const data = new URLSearchParams({
        ticketID: id,
      });

      this.axios
        .post(this.$backendApi + "shop/get-ticket",data)
        .then((response) => {
          console.log(response.data);
          this.ticket = response.data;
        });
    },
  },
  mounted() {
    this.getTicket(sessionStorage.getItem("ticketID"))
    console.log(this.ticket)

    this.socket = new WebSocket("wss://axzrwmh7sb.execute-api.us-east-1.amazonaws.com/production?ticketID="+sessionStorage.getItem("ticketID"));
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
};
</script>