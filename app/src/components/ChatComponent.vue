<template>
  <q-page class="q-pa-md">
    <q-card class="chat-container">
      <q-card-section class="chat-header">
        <div class="text-h6">{{ chatName }}</div>
      </q-card-section>

      <q-separator />

      <q-card-section class="scroll-area">
        <div v-for="(message, index) in messages" :key="index" class="message">
          <q-avatar :label="message.user.charAt(0).toUpperCase()" />
          <div class="message-content">
            <q-chat-message
              :name="message.user"
              :text="[message.text]"
              :sent="message.sent"
            />
          </div>
        </div>
      </q-card-section>

      <q-separator />

      <q-card-actions class="chat-input">
        <q-input
          bottom-slots
          v-model="newMessage"
          label="Message"
          class="full-width-input"
        >
          <template v-slot:append>
            <q-icon
              v-if="newMessage !== ''"
              name="close"
              @click="newMessage = ''"
              class="cursor-pointer"
            />
          </template>

          <template v-slot:after>
            <q-btn round dense flat @click="sendMessage" icon="send" />
          </template>
        </q-input>
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';

interface Message {
  user: string;
  text: string;
  sent: boolean;
}

const messages = ref<Message[]>([]);
const newMessage = ref('');
const chatName = ref('');

chatName.value = 'Chat Name';

function sendMessage() {
  if (newMessage.value.trim() !== '') {
    const message: Message = {
      user: 'You',
      text: newMessage.value,
      sent: true,
    };
    messages.value.push(message);
    newMessage.value = '';
  }
}
</script>

<style scoped>
.chat-container {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 86px);
}
.chat-header {
  flex-shrink: 0;
}

.scroll-area {
  flex: 1;
  overflow-y: auto;
}

.message {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.message-content {
  margin-left: 10px;
  width: 100%;
  max-width: 400px;
}

.chat-input {
  width: 100%;
  position: relative;
  bottom: 0px;
}

.full-width-input {
  width: 100%;
}
</style>
