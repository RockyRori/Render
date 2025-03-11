import axios from 'axios';

const API_BASE_URL = 'http://localhost:5001';

export async function sendMessage(message: string): Promise<string> {
    try {
        const response = await axios.post<{ response: string }>(`${API_BASE_URL}/chat`, { message });
        return response.data.response;
    } catch (error) {
        console.error('Error sending message:', error);
        return 'Error communicating with chatbot';
    }
}