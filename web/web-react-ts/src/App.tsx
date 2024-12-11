import { useState } from 'react'
import './App.css'
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import {GreetService} from "@buf/dilzio_bsrdemo.connectrpc_es/greet_connect";
const transport = createConnectTransport({
    baseUrl: "http://localhost:8080",
});

const client = createClient(GreetService, transport);

function App() {
    const [inputValue, setInputValue] = useState("");
    const [messages, setMessages] = useState<
        {
            fromMe: boolean;
            message: string;
        }[]
    >([]);
    return <>
        <ol>
            {messages.map((msg, index) => (
                <li key={index}>
                    {`${msg.fromMe ? "ME:" : "Greet Service:"} ${msg.message}`}
                </li>
            ))}
        </ol>
        <form onSubmit={async (e) => {
            e.preventDefault();
            // Clear inputValue since the user has submitted.
            setInputValue("");
            // Store the inputValue in the chain of messages and
            // mark this message as coming from "me"
            setMessages((prev) => [
                ...prev,
                {
                    fromMe: true,
                    message: inputValue,
                },
            ]);
            const response = await client.greet({
                name: inputValue,
            });
            setMessages((prev) => [
                ...prev,
                {
                    fromMe: false,
                    message: response.greeting,
                },
            ]);
        }}>
            <input value={inputValue} onChange={e => setInputValue(e.target.value)} />
            <button type="submit">Send</button>
        </form>
    </>;
}

export default App
