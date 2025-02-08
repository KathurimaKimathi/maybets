CREATE TABLE IF NOT EXISTS bets (
    id TEXT PRIMARY KEY,                    
    bet_id TEXT UNIQUE NOT NULL,             
    user_id TEXT NOT NULL,                  
    amount REAL NOT NULL,                  
    odds REAL NOT NULL,                    
    outcome TEXT CHECK(outcome IN ('win', 'lose')) NOT NULL,  
    timestamp TEXT NOT NULL,                
    created TEXT NOT NULL,                  
    updated TEXT NOT NULL,                  
    created_by TEXT,                        
    updated_by TEXT                      
);

CREATE INDEX IF NOT EXISTS idx_user_id ON bets(user_id);