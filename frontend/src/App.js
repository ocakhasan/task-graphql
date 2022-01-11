import React, {useState} from 'react';
import 'graphiql/graphiql.css';
import { GraphiQL } from 'graphiql';
import './App.css'

const App = () => {
  let [heroes, setHeroes] = useState()
  let [show, setShow] = useState(true)

  const URL = 'http://localhost:8085/query'

  const fetcher =  async graphQLParams => {
    try {
      const data = await fetch(
        URL,
        {
          method: 'POST',
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(graphQLParams),
          credentials: 'same-origin',
        },
      );
      
      let resp = await data.json()
      if (resp.data.heroes) {
        setHeroes(resp.data.heroes)
      }

      return resp
    } catch(error) {
      console.log(error)
    }
  }

  return (
    <div className='container'>
      <button className="toggle" onClick={() => setShow(!show)}>Toggle </button>
      <div className='heroes-div'>
        {heroes && heroes.map(hero => (
          <div className='hero-div'>
            <p>{hero.name}</p>
          </div>
        ))
        }
      </div>
      {show? <div className="graphiql-div">
        <GraphiQL
          fetcher={fetcher} 
        >
        </GraphiQL>
      </div>: <></>}
  </div>
  )
}

export default App;
