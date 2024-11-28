import React from 'react';
import './App.css';
import logo from './assets/images/RDOLogo.png';
import A196 from './assets/images/A196.jpg';
import A202 from './assets/images/A202.jpg';
import { WindowMinimise, Quit } from '../wailsjs/runtime/runtime';
import { LaunchRedM } from '../wailsjs/go/main/App';

declare global {
    interface Window {
        go: {
            main: {
                App: {
                    LaunchRedM(): void;
                }
            }
        }
    }
}

function App() {
    const handleCardClick = (url:string) => {
        window.open(url, '_blank');
    };

    return (
        <div className="app" id="App">
            {/* Barre personnalisée */}
            <header style={{widows: 1}}>
            <div className="custom-titlebar">
                <div className="window-controls">
                    <button className="window-button" onClick={WindowMinimise}>−</button>
                    <button className="window-button close" onClick={Quit}>×</button>
                </div>
            </div>
            </header>

            {/* Header */}
            <div className="header">
                <button className="header-button" onClick={() => window.open('https://example.com', '_blank')}>JEUX</button>
                <button className="header-button" onClick={() => window.open('https://example.com', '_blank')}>BOUTIQUE</button>
                <button className="header-button" onClick={() => window.open('https://example.com', '_blank')}>PARAMÈTRES</button>
            </div>

            {/* Container gauche */}
            <div className="container-left">
                <img src={logo} id="logo" alt="logo" />
                <button 
                    className="button-play" 
                    onClick={() => {
                        console.log("Bouton cliqué");
                        LaunchRedM();
                    }}
                >
                    JOUER À REDM
                </button>
            </div>

            {/* Container droit */}
            <div className="container-right">
                <div className="card" onClick={() => handleCardClick('https://example.com/card1')}>
                    <img src={A196} alt="Card 1" />
                    <p>Double Rewards for Skilled Thieves on Camp and Homestead Robberies</p>
                </div>
                <div className="card" onClick={() => handleCardClick('https://example.com/card2')}>
                    <img src={A202} alt="Card 2" />
                    <p>Bonus dans Route commerciale</p>
                </div>
            </div>
        </div>
    );
}

export default App;
