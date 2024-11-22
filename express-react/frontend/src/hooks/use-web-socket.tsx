import { useEffect } from 'react';
import { WebSocketService } from '../service/websocket';

// Singleton instance initialized with env variable
const webSocketService = new WebSocketService();

export const useWebSocket = () => {
  useEffect(() => {
    // Connect when component mounts
    webSocketService.connect();

    // Disconnect when component unmounts
    return () => {
      webSocketService.disconnect();
    };
  }, []);

  // Return the service instance for use in components
  return webSocketService;
};