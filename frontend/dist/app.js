let isRunning = false;
let selectedButton = 'left';
let selectedType = 'single';

// Initialize segmented controls
function initSegmentedControl(containerId) {
    const container = document.getElementById(containerId);
    const buttons = container.querySelectorAll('button');
    
    buttons.forEach(button => {
        button.addEventListener('click', () => {
            buttons.forEach(btn => btn.classList.remove('active'));
            button.classList.add('active');
            
            if (containerId === 'buttonControl') {
                selectedButton = button.dataset.value;
            } else if (containerId === 'typeControl') {
                selectedType = button.dataset.value;
            }
        });
    });
}

// Update status UI
function updateStatus(running) {
    isRunning = running;
    const statusDot = document.getElementById('statusDot');
    const statusText = document.getElementById('statusText');
    const toggleBtn = document.getElementById('toggleBtn');
    
    if (running) {
        statusDot.classList.add('running');
        statusText.classList.add('running');
        statusText.textContent = 'Active';
        toggleBtn.innerHTML = '<i class="fa-solid fa-stop"></i> Stop';
        toggleBtn.classList.remove('btn-primary');
        toggleBtn.classList.add('btn-danger');
    } else {
        statusDot.classList.remove('running');
        statusText.classList.remove('running');
        statusText.textContent = 'Idle';
        toggleBtn.innerHTML = '<i class="fa-solid fa-play"></i> Start';
        toggleBtn.classList.remove('btn-danger');
        toggleBtn.classList.add('btn-primary');
    }
}

// Toggle clicking
async function toggleClicking() {
    try {
        // Check actual state from Go backend
        const actuallyRunning = await window.go.main.App.IsRunning();
        
        if (actuallyRunning) {
            await window.go.main.App.StopClicking();
            updateStatus(false);
        } else {
            const interval = parseFloat(document.getElementById('interval').value);
            
            if (interval <= 0 || isNaN(interval)) {
                alert('Please enter a valid interval greater than 0');
                return;
            }
            
            const config = {
                interval: interval,
                button: selectedButton,
                clickType: selectedType
            };
            
            await window.go.main.App.StartClicking(config);
            updateStatus(true);
        }
    } catch (error) {
        alert('Error: ' + error);
        console.error('Toggle error:', error);
    }
}

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    initSegmentedControl('buttonControl');
    initSegmentedControl('typeControl');
    
    document.getElementById('toggleBtn').addEventListener('click', toggleClicking);
    
    // Listen for status changes from Go (including global hotkey)
    if (window.runtime) {
        window.runtime.EventsOn('status_changed', (status) => {
            updateStatus(status === 'running');
        });
    }
    
    // Sync UI state with Go backend on startup
    if (window.go && window.go.main && window.go.main.App) {
        window.go.main.App.IsRunning().then(running => {
            updateStatus(running);
        }).catch(err => {
            console.error('Failed to get initial state:', err);
        });
    }
});
