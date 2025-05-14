package main

import (
	"fmt"
	"log"
	"net/http"
)

const htmlContent = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to AwesomeApp!</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap" rel="stylesheet">
    <style>
        /* Apply Inter font to the body */
        body { 
            font-family: 'Inter', sans-serif; 
        }
        /* Custom gradient for the hero section */
        .hero-gradient { 
            background-image: linear-gradient(to right, #6366F1, #A855F7); /* Indigo to Purple */
        }
        /* Subtle transition and hover effects for CTA buttons */
        .cta-button { 
            transition: all 0.3s ease; 
        }
        .cta-button:hover { 
            transform: translateY(-2px); 
            box-shadow: 0 10px 15px -3px rgba(0,0,0,0.1), 0 4px 6px -2px rgba(0,0,0,0.05); 
        }
        /* Hover effects for feature cards */
        .feature-card { 
            transition: transform 0.3s ease, box-shadow 0.3s ease; 
        }
        .feature-card:hover { 
            transform: translateY(-5px); 
            box-shadow: 0 20px 25px -5px rgba(0,0,0,0.1), 0 10px 10px -5px rgba(0,0,0,0.04); 
        }
        /* Ensure content below fixed navbar is visible */
        body {
            padding-top: 4rem; /* Height of the navbar (h-16) */
        }
    </style>
</head>
<body class="bg-slate-50 text-slate-800 antialiased">

    <nav class="bg-white/80 backdrop-blur-md shadow-sm fixed w-full z-50 top-0 left-0 right-0">
        <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex items-center justify-between h-16">
                <div class="flex items-center">
                    <a href="#" class="text-2xl font-bold text-indigo-600">AwesomeApp</a>
                </div>
                <div class="hidden md:block">
                    <div class="ml-10 flex items-baseline space-x-4">
                        <a href="#features" class="text-slate-600 hover:text-indigo-600 px-3 py-2 rounded-md text-sm font-medium">Features</a>
                        <a href="#pricing" class="text-slate-600 hover:text-indigo-600 px-3 py-2 rounded-md text-sm font-medium">Pricing</a>
                        <a href="#" class="bg-indigo-600 hover:bg-indigo-700 text-white px-3 py-2 rounded-md text-sm font-medium cta-button">Sign Up</a>
                    </div>
                </div>
                <div class="-mr-2 flex md:hidden">
                    <button type="button" id="mobile-menu-button" class="bg-slate-100 inline-flex items-center justify-center p-2 rounded-md text-slate-500 hover:text-indigo-600 hover:bg-slate-200 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500" aria-controls="mobile-menu" aria-expanded="false">
                        <span class="sr-only">Open main menu</span>
                        <svg class="block h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
                        </svg>
                        <svg class="hidden h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                        </svg>
                    </button>
                </div>
            </div>
        </div>
        <div class="md:hidden hidden" id="mobile-menu">
            <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
                <a href="#features" class="text-slate-600 hover:bg-slate-200 hover:text-indigo-600 block px-3 py-2 rounded-md text-base font-medium">Features</a>
                <a href="#pricing" class="text-slate-600 hover:bg-slate-200 hover:text-indigo-600 block px-3 py-2 rounded-md text-base font-medium">Pricing</a>
            </div>
            <div class="pt-4 pb-3 border-t border-slate-200">
                <div class="px-5">
                     <a href="#" class="w-full block text-center bg-indigo-600 hover:bg-indigo-700 text-white px-3 py-2 rounded-md text-base font-medium cta-button">Sign Up</a>
                </div>
            </div>
        </div>
    </nav>

    <header class="hero-gradient pt-20 pb-20 text-white"> {/* Adjusted pt-20 from pt-32 as body now has padding-top */}
        <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
            <h1 class="text-4xl sm:text-5xl md:text-6xl font-extrabold leading-tight">
                Discover Something <span class="block">Truly Amazing</span>
            </h1>
            <p class="mt-6 text-lg sm:text-xl text-indigo-100 max-w-2xl mx-auto">
                Our revolutionary platform empowers you to achieve your goals faster, smarter, and with more style than ever before. Join the future, today.
            </p>
            <div class="mt-10 flex flex-col sm:flex-row sm:justify-center space-y-4 sm:space-y-0 sm:space-x-4">
                <a href="#" class="bg-white text-indigo-600 hover:bg-indigo-50 text-lg font-semibold px-8 py-3 rounded-lg shadow-lg cta-button inline-block">
                    Get Started Now
                </a>
                <a href="#features" class="text-white hover:text-indigo-200 border border-white/50 hover:border-white/80 text-lg font-semibold px-8 py-3 rounded-lg inline-block cta-button">
                    Learn More &rarr;
                </a>
            </div>
        </div>
    </header>

    <section id="features" class="py-20 bg-white">
        <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="text-center mb-16">
                <h2 class="text-3xl sm:text-4xl font-bold text-slate-900">Why Choose AwesomeApp?</h2>
                <p class="mt-4 text-lg text-slate-600">We're not just another app. We're a solution.</p>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
                <div class="bg-slate-50 p-8 rounded-xl shadow-lg feature-card">
                    <div class="flex items-center justify-center h-12 w-12 rounded-full bg-indigo-500 text-white mb-6">
                        <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" /></svg>
                    </div>
                    <h3 class="text-xl font-semibold text-slate-900 mb-3">Innovative Solutions</h3>
                    <p class="text-slate-600">Cutting-edge technology designed to solve real-world problems efficiently and effectively.</p>
                </div>
                <div class="bg-slate-50 p-8 rounded-xl shadow-lg feature-card">
                    <div class="flex items-center justify-center h-12 w-12 rounded-full bg-purple-500 text-white mb-6">
                         <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                    </div>
                    <h3 class="text-xl font-semibold text-slate-900 mb-3">User-Friendly Interface</h3>
                    <p class="text-slate-600">Intuitive design that makes navigation and usage a breeze for everyone, regardless of tech-savviness.</p>
                </div>
                <div class="bg-slate-50 p-8 rounded-xl shadow-lg feature-card">
                     <div class="flex items-center justify-center h-12 w-12 rounded-full bg-pink-500 text-white mb-6">
                        <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" /></svg>
                    </div>
                    <h3 class="text-xl font-semibold text-slate-900 mb-3">Dedicated Support</h3>
                    <p class="text-slate-600">Our expert team is always ready to assist you, ensuring a smooth and productive experience.</p>
                </div>
            </div>
        </div>
    </section>

    <section id="pricing" class="py-20 bg-slate-100">
        <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="text-center mb-16">
                <h2 class="text-3xl sm:text-4xl font-bold text-slate-900">Simple, Transparent Pricing</h2>
                <p class="mt-4 text-lg text-slate-600">Choose the plan that's right for you.</p>
            </div>
            <div class="text-center">
                <p class="text-slate-700 text-xl">Pricing details coming soon! Stay tuned.</p>
                <a href="#" class="mt-8 bg-indigo-600 hover:bg-indigo-700 text-white text-lg font-semibold px-8 py-3 rounded-lg shadow-md cta-button inline-block">
                    Notify Me
                </a>
            </div>
        </div>
    </section>

    <footer class="bg-slate-800 text-slate-300 py-12">
        <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
            <p>&copy; 2025 AwesomeApp. All rights reserved.</p>
            <p class="mt-2 text-sm">Built with <span class="text-red-500">&hearts;</span> and Go.</p>
        </div>
    </footer>

    <script>
        // Mobile menu toggle
        const mobileMenuButton = document.getElementById('mobile-menu-button');
        const mobileMenu = document.getElementById('mobile-menu');
        // Get both SVGs inside the button
        const svgs = mobileMenuButton.querySelectorAll('svg'); 

        mobileMenuButton.addEventListener('click', () => {
            mobileMenu.classList.toggle('hidden');
            // Toggle visibility of burger and close icons
            svgs[0].classList.toggle('hidden'); 
            svgs[0].classList.toggle('block');
            svgs[1].classList.toggle('hidden');
            svgs[1].classList.toggle('block');
        });

        // Smooth scroll for internal links
        document.querySelectorAll('a[href^="#"]').forEach(anchor => {
            anchor.addEventListener('click', function (e) {
                e.preventDefault();
                const targetId = this.getAttribute('href');
                // Don't scroll for href="#" as it's used for placeholder links
                if (targetId === '#') return; 

                const targetElement = document.querySelector(targetId);
                if (targetElement) {
                    // Account for fixed navbar height
                    const navbarHeight = document.querySelector('nav').offsetHeight;
                    const targetPosition = targetElement.getBoundingClientRect().top + window.pageYOffset - navbarHeight;
                    
                    window.scrollTo({
                        top: targetPosition,
                        behavior: 'smooth'
                    });
                }
            });
        });
    </script>

</body>
</html>
`

// Handles equest to "/"
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlContent)
}

// Handles request to "/health"
func handlerHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status": "ok"}`)
}

func main() {
	port := ":8900"
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/health", handlerHealth)
	log.Printf("Server starting on http://localhost:%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
